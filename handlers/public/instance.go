package public

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/config"
	"net/http"
)

func Instance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"link": config.Config.Misskey.Instance,
	})
}
