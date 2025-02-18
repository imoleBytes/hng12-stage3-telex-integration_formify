package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	// telex will access these
	r.GET("/formify-integration.json", HandleIntegrationJSON)
	r.POST("/generate-formify", HandleGenerate)

	// this to be accessed through the html form
	r.POST("/formify/website/:id", HandleFormSubmission)

	// the index route for health
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "server running fine",
		})
	})

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
