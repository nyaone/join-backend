package admin

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/global"
	"join-nyaone/models"
	"net/http"
)

func InviteeCount(ctx *gin.Context) {
	userId, exist := ctx.Get("UID")
	if !exist {
		global.Logger.Errorf("No UID provided by middleware")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Middleware not working properly",
		})
		return
	}

	var inviteesCount int64
	global.DB.Model(&models.User{}).Where("invited_by_user_id = ?", userId.(uint)).Count(&inviteesCount)

	ctx.JSON(http.StatusOK, gin.H{
		"count": inviteesCount,
	})

}
