package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	firebase "firebase.google.com/go"

	"github.com/gin-contrib/cors"
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
	_, err := collection.DeleteMany(ctx, filter)
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
	log.Println(db_uri, db_username, db_password, db_host)

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

func main() {
	godotenv.Load(".env")
	client, err := loadMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	app.Auth(context.Background())

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/slugs", func(ctx *gin.Context) {
		id := ctx.Query("userid")
		slugs, err := getSlugs(id, client)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, gin.H{})
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
			log.Fatal(err)
			ctx.JSON(400, err)
		}
		log.Println(body)
		var slug = &Slug{
			ID:     primitive.NewObjectID(), // use client to gen uuid
			UserID: body.UserID,
			Slug:   body.Slug,
			Domain: body.Domain,
		}

		err = createSlug(slug, client)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, err)
		}

		ctx.JSON(200, slug)
	})
	r.DELETE("/slugs", func(ctx *gin.Context) {
		uid := ctx.Query("userid")
		id_str := ctx.Query("id")
		id, err := primitive.ObjectIDFromHex(id_str)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, err)
		}
		err = deleteSlug(id, uid, client)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, err)
		}

		ctx.JSON(200, bson.D{})
	})
	r.PUT("/slugs", func(ctx *gin.Context) {
		slug := &Slug{}
		err := ctx.ShouldBindBodyWith(slug, binding.JSON)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(400, err)
		}
		err = updateSlug(slug, client)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, err)
		}
		ctx.JSON(200, bson.D{})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
