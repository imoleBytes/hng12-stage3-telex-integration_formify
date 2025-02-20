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

	println(website, "and", channel_id)

	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Unable to parse form"})
		return
	}
	// Retrieve all form data dynamically
	formData := make(map[string]interface{})
	for key, values := range ctx.Request.PostForm {
		if len(values) == 1 {
			formData[key] = values[0] // Store as string if single value
		} else {
			formData[key] = values // Store as slice if multiple values
		}
	}

	tmpl := FormatFormDataToHTML(formData)

	// formDatastr, _ := json.Marshal(formData)

	err = WebhookSendData(tmpl, channel_id, website)
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

func FormatFormDataToHTML(form map[string]interface{}) string {
	msg := "*******************\n"

	for key, val := range form {
		msg += fmt.Sprintf("%s:\t\t%v\n", key, val)
	}
	msg += "*******************\n"
	return msg
}

func WebhookSendData(formdata, channel_id, website string) error {
	url := "https://ping.telex.im/v1/webhooks/" + channel_id

	data := map[string]string{
		"event_name": website + ": form submission",
		"message":    formdata,
		"status":     "success",
		"username":   website + ": form submission",
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
