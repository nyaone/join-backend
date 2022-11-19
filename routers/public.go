package routers

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/handlers/public"
)

func Public(rg *gin.RouterGroup) {
	rg.GET("/healthcheck", public.HealthCheck)
}
