package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	app.Auth(context.Background())
	r := gin.Default()
	r.GET("/slugs", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
