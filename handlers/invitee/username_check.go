package invitee

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/global"
	"join-nyaone/misskey"
	"net/http"
)

func UsernameCheck(ctx *gin.Context) {
	username := ctx.Param("username")
	valid, err := misskey.CheckUsername(username)
	if err != nil {
		global.Logger.Errorf("Failed to check username status with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"valid": valid,
		})
	}

}
