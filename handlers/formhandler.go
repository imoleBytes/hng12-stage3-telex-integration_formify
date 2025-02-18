package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFormSubmission(ctx *gin.Context) {
	website := ctx.Param("website")
	channel_id := ctx.Param("channel_id")

	err := WebhookSendData(channel_id, website)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, map[string]string{
		// "website": website,
		"id":      channel_id,
		"website": website,
		"message": "form submitted successfully!!!!!!",
	})
}

func WebhookSendData(channel_id, website string) error {
	url := "https://ping.telex.im/v1/webhooks/" + channel_id
	data := map[string]string{
		"event_name": "string",
		"message":    "form submitted!!!!",
		"status":     "success",
		"username":   website,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response:", resp.Status)
	return nil
}
