package routers

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/handlers/invitee"
)

func Invitee(rg *gin.RouterGroup) {
	rg.POST("/code/check/:code", invitee.CodeCheck)
	rg.POST("/username/check/:username", invitee.UsernameCheck)
	rg.POST("/register", invitee.Register)
}
