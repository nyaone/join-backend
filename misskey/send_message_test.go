package misskey

import (
	"go.uber.org/zap"
	"join-nyaone/config"
	"join-nyaone/global"
	"testing"
)

func TestSendMessage(t *testing.T) {
	// Prepare
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	config.Config.Misskey.Instance = "https://mk.nyawork.dev"
	config.Config.Misskey.Token.Notify = ""
	targetUserID := "8zmmyb1jkl"
	text := "**Click** ?[here](https://nya.one) to NyaOne\n\nOr ?[here](https://docs.nya.one) to docs?"

	t.Log(SendMessage(targetUserID, text))

}
