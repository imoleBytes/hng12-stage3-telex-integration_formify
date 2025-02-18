package handlers

import "github.com/gin-gonic/gin"

func HandleFormSubmission(ctx *gin.Context) {
	website := ctx.Param("website")
	channel_id := ctx.Param("channel_id")

	ctx.JSON(200, map[string]string{
		// "website": website,
		"id":      channel_id,
		"website": website,
	})
}
