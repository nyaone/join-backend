package routers

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/handlers/login"
)

func Login(rg *gin.RouterGroup) {
	rg.POST("/request/:username", login.Request)
	rg.POST("/confirm/:token", login.Confirm)
}
