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
	var msgReq MsgRequest
	// var data map[string]interface{}

	err := ctx.Bind(&msgReq)
	if err != nil {
		log.Println("error bindind post data: ", err)
		ctx.JSON(400, gin.H{
			"error": " invalid JSON payload",
		})
		return
	}

	log.Printf("All request data from telex: %+v\n", msgReq)
	fmt.Println("*******************")
	log.Printf("Channel ID is: %s\n", msgReq.ChannelID)
	log.Printf("Message is: %s\n", msgReq.Message)
	log.Printf("Settings are: %+v\n", msgReq.Settings)
	fmt.Println("*******************")

	text := ExtractText(msgReq.Message)
	if text != "/generate_url" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "invalid command",
		})
		return
	}
	log.Println("command is ", text)

	form_name := msgReq.Settings[0].Default

	url := GenerateUniqueURL(msgReq.Settings)

	msg := fmt.Sprintf("Here's the url for <b>[%s Form]</b>: , (%s)\n........................\n", form_name, url)
	msg += "Put the url in the action attribute of your form and set the method to POST\n"
	msg += "Sit back and start getting data from the form in this channel!...\n"

	ctx.JSON(200, gin.H{
		"event_name": "Unique URL Generated",
		// "message":    "Hello, Use this Url for the form world. I hope you are happy happy today",
		"message":  msg,
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
