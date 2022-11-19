package middlewares

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"join-nyaone/consts"
	"join-nyaone/global"
	"net/http"
)

func AuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		global.Logger.Debugf("Checking Authorization token...")

		sess := ctx.GetHeader("Authorization")
		if sess == "" {
			// No session found
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "No session provided",
			})
			ctx.Abort()
			return
		}

		sessionKey := fmt.Sprintf(consts.REDIS_KEY_VALID_SESSION, sess)
		if exist, err := global.Redis.Exists(context.Background(), sessionKey).Result(); err != nil {
			global.Logger.Errorf("Failed to check session key (%s) from redis with error: %v", sessionKey, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to check session from redis: %s", err.Error()),
			})
			ctx.Abort()
			return
		} else if exist == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Session doesn't exist.",
			})
			ctx.Abort()
			return
		}

		userId, err := global.Redis.Get(context.Background(), sessionKey).Uint64()
		if err != nil {
			global.Logger.Errorf("Failed to get session key (%s) from redis with error: %v", sessionKey, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to get session from redis: %s", err.Error()),
			})
			ctx.Abort()
			return
		}

		// Continue
		ctx.Set("UID", uint(userId))
		//ctx.Next()
	}
}
