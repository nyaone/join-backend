package models

import (
	"gorm.io/gorm"
	"join-nyaone/types"
	"time"
)

type Code struct {
	// Database meta
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	types.Code
}
