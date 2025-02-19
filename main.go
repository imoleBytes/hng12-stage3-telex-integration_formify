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

	// Telex will access these
	// ************for getting the integration's json**********
	r.GET("/formify-integration.json", handlers.HandleIntegrationJSON)

	// *********** this handles the integration commands. in this case , the `/generate_url` command.*************
	r.POST("/generate-formify", handlers.HandleGenerate)

	// The endusers access the integration through this
	// *********** this to be accessed through the html form ************
	r.POST("/formify/:website/:channel_id", handlers.HandleFormSubmission)

	// the index route for health
	r.GET("/", handlers.HandleIndex)

	if err := r.Run(); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
