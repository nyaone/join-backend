package inits

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"join-nyaone/config"
	"join-nyaone/utils"
	"os"
)

func Config() error {
	// Read config file
	configFileBytes, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFileBytes, &config.Config)
	if err != nil {
		return err
	}

	// Validate config
	if !utils.ValidateUri(config.Config.FrontendUri) {
		return fmt.Errorf("invalid frontend uri")
	}
	if !utils.ValidateUri(config.Config.Misskey.Instance) {
		return fmt.Errorf("invalid misskey instance uri")
	}
	if config.Config.Misskey.Token.Admin == "" {
		return fmt.Errorf("misskey admin token is empty")
	}
	if config.Config.Misskey.Token.Notify == "" {
		return fmt.Errorf("misskey notify token is empty")
	}

	return nil
}
