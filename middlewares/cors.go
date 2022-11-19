package middlewares

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/config"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if config.Config.System.Debug || ctx.GetHeader("Origin") == config.Config.FrontendUri {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(http.StatusNoContent)
			}
		} else {
			// Otherwise block requests for safety concern
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
