package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var IntegrationJSON = `{
	"data": {
		"date": {
		"created_at": "2025-02-17",
		"updated_at": "2025-02-17"
		},
		"descriptions": {
		"app_name": "Formify FaaS",
		"app_description": "Formify is a seamless Form as a Service (FaaS) tool that empowers users to collect and manage form data without the need for a backend infrastructure. With Formify, users can simply embed a URL into the 'action' attribute of an HTML form, and the submitted data will be automatically processed and sent to a designated Telex channel.\n\nGenerate unique action attribbute value for your HTML Forms",
		"app_logo": "https://img.freepik.com/free-vector/retro-circular-pattern-design_1308-175051.jpg?t=st=1739808712~exp=1739812312~hmac=f03b43859fc31fbdd675f1907599bd626959e488186958f208d879c6fd1ef10a&w=740",

		"app_url": "http://100.25.134.239",
		"background_color": "#fff"
		},
		"is_active": true,
		"integration_type": "modifier",
		"key_features": [
		"No Backend Required",
		"Easy Integration",
		"Real-time Data Submission",
		"Scalable and Secure",
		"No Extra Coding Required"
		],
		"author": "Imoleayo Kolawole",
		"settings": [
		{
			"label": "Form Name",
			"type": "text",
			"required": true,
			"default": ""
		},
		{
			"label": "Website",
			"type": "text",
			"required": true,
			"default": ""
		}
		],
		"target_url": "http://100.25.134.239/generate-formify",		
	}
}`

// this returns the integration json
func HandleIntegrationJSON(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte(IntegrationJSON))
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
