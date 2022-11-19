package models

import (
	"gorm.io/gorm"
	"join-nyaone/types"
)

type User struct {
	// Database meta
	gorm.Model

	types.User
}
