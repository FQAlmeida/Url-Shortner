package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Slug struct {
	ID     primitive.ObjectID `json:"id"`
	Slug   string             `json:"slug"`
	Domain string             `json:"redirect"`
	UserID string             `json:"uid"`
}
type SlugCreate struct {
	Slug   string `json:"slug"`
	Domain string `json:"redirect"`
	UserID string `json:"uid"`
}

func createSlug(slug *Slug, client *mongo.Client) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("url-shortner").Collection("slugs")
	_, err := collection.InsertOne(ctx, slug)
	return err
}

func deleteSlug(slug_id primitive.ObjectID, uid string, client *mongo.Client) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("url-shortner").Collection("slugs")
	filter := bson.D{{
		Key:   "userid",
		Value: uid,
	}, {
		Key:   "id",
		Value: slug_id,
	}}
	log.Println(filter)
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
func updateSlug(slug *Slug, client *mongo.Client) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("url-shortner").Collection("slugs")
	filter := bson.D{{
		Key:   "userid",
		Value: slug.UserID,
	}, {
		Key:   "id",
		Value: slug.ID,
	}}
	updater := bson.D{{Key: "$set", Value: bson.D{{
		Key:   "slug",
		Value: slug.Slug,
	}, {
		Key:   "domain",
		Value: slug.Domain,
	}}}}
	_, err := collection.UpdateMany(ctx, filter, updater)
	return err
}

func getSlugs(uid string, client *mongo.Client) ([]*Slug, error) {
	filter := bson.D{{
		Key:   "userid",
		Value: uid,
	}}
	return filterTasks(filter, client)
}
func getSlug(slug string, client *mongo.Client) (*Slug, error) {
	filter := bson.D{{
		Key:   "slug",
		Value: slug,
	}}
	slugs, err := filterTasks(filter, client)
	if len(slugs) == 0 {
		return nil, errors.New("Slug not found")
	}
	return slugs[len(slugs)-1], err
}

func filterTasks(filter interface{}, client *mongo.Client) ([]*Slug, error) {
	// A slice of slugs for storing the decoded documents
	var slugs []*Slug

	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("url-shortner").Collection("slugs")

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return slugs, err
	}
	for cur.Next(ctx) {
		var t Slug
		err := cur.Decode(&t)
		if err != nil {
			return slugs, err
		}

		slugs = append(slugs, &t)
	}

	if err := cur.Err(); err != nil {
		return slugs, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(slugs) == 0 {
		slugs := []*Slug{}
		return slugs, nil
	}
	return slugs, nil
}

func loadMongoClient() (*mongo.Client, error) {
	db_username := url.QueryEscape(os.Getenv("DATABASE_USERNAME"))
	db_password := url.QueryEscape(os.Getenv("DATABASE_PASSWORD"))
	db_host := url.QueryEscape(os.Getenv("DATABASE_HOST"))

	db_uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", db_username, db_password, db_host,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(
			db_uri))
	if err != nil {
		log.Fatal(err)
		return client, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func loadFirebaseClient() (*auth.Client, error) {
	// firebase_api_key := os.Getenv("FIREBASE_API_KEY")
	// firebase_app_id := os.Getenv("FIREBASE_APP_ID")
	// firebase_auth_domain := os.Getenv("FIREBASE_AUTH_DOMAIN")

	opts := option.WithCredentialsFile("url-shortner-fqa-firebase-adminsdk-zmv68-a3dcfdbe88.json")

	app, err := firebase.NewApp(context.Background(), nil, opts)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	fireAuth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v\n", err)
	}
	return fireAuth, err
}

func checkUserExists(userid string, firebaseAuth *auth.Client) (bool, error) {
	result, err := firebaseAuth.GetUser(context.Background(), userid)
	if auth.IsUserNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return result != nil, err
}

func main() {
	godotenv.Load(".env")
	client, err := loadMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	firebaseAuth, err := loadFirebaseClient()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/slugs", func(ctx *gin.Context) {
		id := ctx.Query("userid")
		exists, err := checkUserExists(id, firebaseAuth)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err})
			return
		}
		if !exists {
			ctx.JSON(400, gin.H{"err": errors.New("USER not found")})
			return
		}
		slugs, err := getSlugs(id, client)
		if err != nil {
			ctx.JSON(500, gin.H{})
			return
		}
		ctx.JSON(200, slugs)
	})
	r.GET("/slug", func(ctx *gin.Context) {
		slug := ctx.Query("slug")
		slug_data, err := getSlug(slug, client)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}
		ctx.JSON(200, slug_data)
	})
	r.POST("/slugs", func(ctx *gin.Context) {
		body := SlugCreate{}
		err := ctx.ShouldBindBodyWith(&body, binding.JSON)
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		exists, err := checkUserExists(body.UserID, firebaseAuth)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err})
			return
		}
		if !exists {
			ctx.JSON(400, gin.H{"err": errors.New("USER not found")})
			return
		}

		var slug = &Slug{
			ID:     primitive.NewObjectID(), // use client to gen uuid
			UserID: body.UserID,
			Slug:   body.Slug,
			Domain: body.Domain,
		}

		err = createSlug(slug, client)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, slug)
	})
	r.DELETE("/slugs", func(ctx *gin.Context) {
		uid := ctx.Query("userid")

		exists, err := checkUserExists(uid, firebaseAuth)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err})
			return
		}
		if !exists {
			ctx.JSON(400, gin.H{"err": errors.New("USER not found")})
			return
		}

		id_str := ctx.Query("id")
		id, err := primitive.ObjectIDFromHex(id_str)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		err = deleteSlug(id, uid, client)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, bson.D{})
	})
	r.PUT("/slugs", func(ctx *gin.Context) {
		slug := &Slug{}
		err := ctx.ShouldBindBodyWith(slug, binding.JSON)
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		exists, err := checkUserExists(slug.UserID, firebaseAuth)
		if err != nil {
			ctx.JSON(500, gin.H{"err": err})
			return
		}
		if !exists {
			ctx.JSON(400, gin.H{"err": errors.New("USER not found")})
			return
		}

		err = updateSlug(slug, client)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, bson.D{})
	})
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	// r.Run() // listen and serve on 0.0.0.0:8080
	log.Fatal(autotls.RunWithContext(ctx, r, "*"))
}
