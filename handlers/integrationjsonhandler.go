package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

var Data = IntegrationStruct{
	Date: struct {
		CreatedAt string "json:\"created_at\""
		UpdatedAt string "json:\"updated_at\""
	}{
		CreatedAt: "2025-02-17",
		UpdatedAt: "2025-02-17",
	},
	Descriptions: struct {
		AppName         string "json:\"app_name\""
		AppDescription  string "json:\"app_description\""
		AppLogo         string "json:\"app_logo\""
		AppURL          string "json:\"app_url\""
		BackgroundColor string "json:\"background_color\""
	}{
		AppName:         "Formify FaaS",
		AppDescription:  "Formify is a seamless Form as a Service (FaaS) tool that empowers users to collect and manage form data without the need for a backend infrastructure. With Formify, users can simply embed a URL into the 'action' attribute of an HTML form, and the submitted data will be automatically processed and sent to a designated Telex channel.\n\nGenerate unique action attribbute value for your HTML Forms",
		AppLogo:         "https://img.freepik.com/free-vector/retro-circular-pattern-design_1308-175051.jpg?t=st=1739808712~exp=1739812312~hmac=f03b43859fc31fbdd675f1907599bd626959e488186958f208d879c6fd1ef10a&w=740",
		AppURL:          "https://hng12-stage3-telex-integration-formify.onrender.com",
		BackgroundColor: "#fff",
	},
	IntegrationCategory: "CRM & Customer Support",
	IsActive:            true,
	IntegrationType:     "modifier",
	KeyFeatures: []string{"No Backend Required",
		"Easy Integration",
		"Real-time Data Submission",
		"Scalable and Secure",
		"No Extra Coding Required",
	},
	Author: "Imoleayo Kolawole",
	Settings: []Setting{
		{Label: "Form Name", Type: "text", Default: "", Required: true},
		{Label: "Website", Type: "text", Default: "", Required: true},
		{Label: "ChannelID", Type: "text", Default: "", Required: true},
	},
	TargetURL: os.Getenv("BASE_URL") + "/generate-formify",
}

// this returns the integration json
func HandleIntegrationJSON(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": Data,
	})
}
