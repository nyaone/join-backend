package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SessionCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
