package misskey

import (
	"go.uber.org/zap"
	"join-nyaone/config"
	"join-nyaone/global"
	"testing"
)

func TestGetUserID(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.Misskey.Instance = "https://mk.nyawork.dev"
	username := "admin"

	t.Log(GetUser(username))
}
