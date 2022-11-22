package misskey

import (
	"go.uber.org/zap"
	"join-nyaone/config"
	"join-nyaone/global"
	"testing"
)

func TestNotesCreate(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.Misskey.Instance = "https://mk.nyawork.dev"
	config.Config.Misskey.Token.Notify = "OtKfASrl08Mx8JwXoO0x8squQfUPx5pA"

	text := "Test note @candinya :nacho_hi:"

	t.Log(NotesCreate(text))

}
