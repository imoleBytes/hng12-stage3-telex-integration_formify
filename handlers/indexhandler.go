package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handles the index route
func HandleIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "https://github.com/imoleBytes/hng12-stage3-telex-integration_formify/blob/main/README.md")
}
