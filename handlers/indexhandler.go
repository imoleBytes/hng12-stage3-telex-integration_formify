package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handles the index route
func HandleIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"status_code": 200,
		"data": struct {
			Integration string
			Description string
			Author      string
			GithubRepo  string
		}{
			Integration: "Formify FaaS",
			Description: "Formify is a seamless Form as a Service (FaaS) tool that empowers users to collect and manage form data without the need for a backend infrastructure. With Formify, users can simply embed a URL into the 'action' attribute of an HTML form, and the submitted data will be automatically processed and sent to a designated Telex channel.\n\nGenerate unique action attribbute value for your HTML Forms",
			Author:      "Imoleayo Kolawole",
			GithubRepo:  "https://github.com/telexintegrations/hng12-stage3-telex-integration_formify",
		},
	})
	// ctx.Redirect(http.StatusTemporaryRedirect, "https://github.com/imoleBytes/hng12-stage3-telex-integration_formify/blob/main/README.md")
}
