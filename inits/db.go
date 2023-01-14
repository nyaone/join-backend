package inits

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"join-nyaone/config"
	"join-nyaone/global"
	"join-nyaone/models"
)

func DB() error {
	var err error
	var gormConfig gorm.Config

	// Set config
	if config.Config.System.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Warn)
	}

	// Connect to database
	global.DB, err = gorm.Open(postgres.Open(config.Config.System.Postgres), &gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Run migrations
	err = mig()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

func mig() error {
	err := global.DB.AutoMigrate(
		&models.User{},
		&models.Code{},
		&models.Application{},
	)
	if err != nil {
		return err
	}

	return nil
}
