package routers

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/handlers/admin"
)

func Admin(rg *gin.RouterGroup) {
	rg.GET("/code/count", admin.CodeCount)
	rg.GET("/code/list", admin.CodeList)
	rg.POST("/code/create", admin.CodeCreate)
	rg.PUT("/code/edit/:code", admin.CodeEdit)
	
	rg.GET("/invitee/count", admin.InviteeCount)
	rg.GET("/invitee/list", admin.InviteeList)
	rg.GET("/invitee/list/:code", admin.InviteeList)

	rg.DELETE("/logout", admin.Logout)
}
