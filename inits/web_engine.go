package inits

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/config"
	"join-nyaone/routers"
)

type Option func(*gin.Engine)

var options []Option

// Include : Register routers
func include(opts ...Option) {
	options = append(options, opts...)
}

func ginInit(middleware ...gin.HandlerFunc) *gin.Engine {

	if !config.Config.System.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	for _, mid := range middleware {
		engine.Use(mid)
	}

	for _, opt := range options {
		opt(engine)
	}

	return engine
}

func WebEngine() *gin.Engine {
	include(routers.R)

	return ginInit()
}
