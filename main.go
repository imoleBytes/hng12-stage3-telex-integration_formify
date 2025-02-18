package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/imoleBytes/hng12-stage3-telex-integration_formify/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	// telex will access these
	r.GET("/formify-integration.json", handlers.HandleIntegrationJSON)
	r.POST("/generate-formify", handlers.HandleGenerate)

	// this to be accessed through the html form
	r.POST("/formify/website/:id", handlers.HandleFormSubmission)

	// the index route for health
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "server running fine",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
