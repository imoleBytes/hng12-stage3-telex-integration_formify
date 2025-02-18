package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type IntegrationStruct struct {
	Date struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"date"`
	Descriptions struct {
		AppName         string `json:"app_name"`
		AppDescription  string `json:"app_description"`
		AppLogo         string `json:"app_logo"`
		AppURL          string `json:"app_url"`
		BackgroundColor string `json:"background_color"`
	} `json:"descriptions"`
	IntegrationCategory string    `json:"integration_category"`
	IsActive            bool      `json:"is_active"`
	IntegrationType     string    `json:"integration_type"`
	KeyFeatures         []string  `json:"key_features"`
	Author              string    `json:"author"`
	Settings            []Setting `json:"settings"`
	TargetURL           string    `json:"target_url"`
}

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
		AppURL:          "http://100.25.134.239",
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
	Author:    "Imoleayo Kolawole",
	Settings:  []Setting{{Label: "Form Name", Type: "text", Default: "", Required: true}, {Label: "Website", Type: "text", Default: "", Required: true}},
	TargetURL: "http://100.25.134.239/generate-formify",
}

// this returns the integration json
func HandleIntegrationJSON(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": Data,
	})
}

// This generate a unique url to be used in an html form

// "channel_id": "0192dd70-cdf1-7e15-8776-4fee4a78405e",
// "settings": [
//   {"label": "maxMessageLength", "type": "number", "default": 30, "required": true},
//   {"label": "repeatWords", "type": "multi-select", "default": "world, happy", "required": true},
//   {"label": "noOfRepetitions", "type": "number", "default": 2, "required": true}
// ],
// "message": "Hello, world. I hope you are happy today"
// }'

type MsgRequest struct {
	ChannelID string    `json:"channel_id"`
	Settings  []Setting `json:"settings"`
	Message   string    `json:"message"`
}
type Setting struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Default  string `json:"default"`
	Required bool   `json:"required"`
}

func ParseSettings(settings []Setting) (form_name, website string) {

	for _, setting := range settings {
		switch setting.Label {
		case "Form Name":
			form_name = setting.Default

		case "Website":
			website = setting.Default
		}
	}
	return
}

func HandleGenerate(ctx *gin.Context) {
	// var msgReq MsgRequest
	var data map[string]interface{}

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Println("error bindind post data: ", err)
		ctx.JSON(400, gin.H{
			"error": " invalid JSON payload",
		})
		return
	}

	log.Println(data)

	channel_id := data["channel_id"].(string)
	// msg := data["message"].(string)
	// setts := data["settings"].([]map[string]interface{})

	url := "http://localhost/formify/website/" + channel_id

	ctx.JSON(200, gin.H{
		"event_name": "Unique URL Generated",
		// "message":    "Hello, Use this Url for the form world. I hope you are happy happy today",
		"message":  fmt.Sprintf("Here's the url for the Form: , (%s)", url),
		"status":   "success",
		"username": "formify-bot",
	})

	// msg := fmt.Sprintf("Here's the url for Form: %s", msgReq.Settings[])

	// ctx.JSON(202, struct {
	// 	event_name string
	// 	message    string
	// 	status     string
	// 	username   string
	// }{
	// 	event_name: "Unique url Generated",
	// 	message:    "Hello, Use this Url for the form world. I hope you are happy happy today",
	// 	status:     "success",
	// 	username:   "formify-bot",
	// })
}

func HandleFormSubmission(ctx *gin.Context) {
	// website := ctx.Param("website")
	id := ctx.Param("id")

	ctx.JSON(200, map[string]string{
		// "website": website,
		"id": id,
	})
}
