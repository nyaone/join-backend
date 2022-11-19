package models

import (
	"gorm.io/gorm"
	"join-nyaone/types"
)

type Code struct {
	// Database meta
	gorm.Model

	types.Code
}
