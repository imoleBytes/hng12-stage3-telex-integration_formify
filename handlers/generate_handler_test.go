package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleGenerate(t *testing.T) {
	router := gin.Default()
	router.POST("/generate-formify", HandleGenerate)

	t.Run("valid input", func(t *testing.T) {
		body := `{
			"channel_id": "0192dd70-cdf1-7e15-8776-4fee4a78405e",
			"settings": [
			   {"label": "Form Name","type": "text","default": "Contact","required": true},
			   {"label": "Website","type": "text","default": "Umegain","required": true},
			   {"label": "ChannelID","type": "text","default": "0192dd70-cdf1-7e15-8776-4fee4a78405e","required": true}
			 ],
			"message": "<p>/generate_url</p><p></p>"
		 }`
		req, _ := http.NewRequest("POST", "/generate-formify", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// assert.JSONEq(t, `{"message": "User created", "user": {"name": "Alice"}}`, w.Body.String())
	})

	t.Run("invalid input", func(t *testing.T) {
		body := `{
			"channel_id": "0192dd70-cdf1-7e15-8776-4fee4a78405e",
			"settings": [
			   {"label": "Form Name","type": "text","default": "Contact","required": true},
			   {"label": "Website","type": "text","default": "Umegain","required": true},
			   {"label": "ChannelID","type": "text","default": "0192dd70-cdf1-7e15-8776-4fee4a78405e","required": true}
			 ],
			"message": "<p>some msg</p><p></p>"
		 }`
		req, _ := http.NewRequest("POST", "/generate-formify", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"event_name": "Invalid Command", "message": "type '/generate_url' to get the unique url for your html forms", "status": "success", "username": "formify-bot"}`, w.Body.String())
	})
}
