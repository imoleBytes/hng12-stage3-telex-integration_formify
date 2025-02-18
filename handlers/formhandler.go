package handlers

import "github.com/gin-gonic/gin"

func HandleFormSubmission(ctx *gin.Context) {
	// website := ctx.Param("website")
	id := ctx.Param("id")

	ctx.JSON(200, map[string]string{
		// "website": website,
		"id": id,
	})
}
