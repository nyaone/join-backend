package inits

import (
	"fmt"
	"join-nyaone/config"
	"join-nyaone/global"
	"join-nyaone/misskey"
)

func Token() error {

	// Check token
	if adminTokenRes, err := misskey.CheckToken(config.Config.Misskey.Token.Admin); err != nil {
		return fmt.Errorf("misskey admin token is not working: %v", err)
	} else if !adminTokenRes.IsAdmin {
		return fmt.Errorf("misskey admin should have admin permission")
	} else {
		global.Logger.Debugf("Admin token initialized, I'm %s (%s)", adminTokenRes.Name, adminTokenRes.Username)
	}

	if notifyTokenRes, err := misskey.CheckToken(config.Config.Misskey.Token.Notify); err != nil {
		return fmt.Errorf("misskey notify token is not working: %v", err)
	} else {
		global.Logger.Debugf("Notify token initialized, I'm %s (%s)", notifyTokenRes.Name, notifyTokenRes.Username)
	}

	return nil
}
