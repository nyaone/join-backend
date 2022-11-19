package routers

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/middlewares"
)

func R(e *gin.Engine) {
	e.Use(middlewares.CORS())

	// Public
	publicApi := e.Group("")
	Public(publicApi)

	// Invitee // !! double e !!
	inviteeApi := e.Group("/invitee")
	Invitee(inviteeApi)

	// Login
	loginApi := e.Group("/login")
	Login(loginApi)

	// Admin
	adminApi := e.Group("/admin")
	adminApi.Use(middlewares.AuthAdmin())
	Admin(adminApi)
}
