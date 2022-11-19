package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"join-nyaone/consts"
	"join-nyaone/global"
	"net/http"
)

func Logout(ctx *gin.Context) {
	sess := ctx.GetHeader("Authorization")
	sessionKey := fmt.Sprintf(consts.REDIS_KEY_VALID_SESSION, sess)
	global.Redis.Del(context.Background(), sessionKey)

	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
