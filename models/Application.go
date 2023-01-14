package models

import (
	"gorm.io/gorm"
	"join-nyaone/types"
)

type Application struct {
	gorm.Model

	// Identifiers
	InstanceUri string `gorm:"index;column:instance_uri"` // Need for instance change
	FrontendUri string `gorm:"index;column:frontend_uri"` // Need for redirect

	// Application Details
	types.ApplicationPublic
	types.ApplicationPrivate
}
