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

	if err := r.Run(); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
