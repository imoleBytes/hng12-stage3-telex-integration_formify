package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// This generate a unique url to be used in an html form

// "channel_id": "0192dd70-cdf1-7e15-8776-4fee4a78405e",
// "settings": [
//   {"label": "maxMessageLength", "type": "number", "default": 30, "required": true},
//   {"label": "repeatWords", "type": "multi-select", "default": "world, happy", "required": true},
//   {"label": "noOfRepetitions", "type": "number", "default": 2, "required": true}
// ],
// "message": "Hello, world. I hope you are happy today"
// }'

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
