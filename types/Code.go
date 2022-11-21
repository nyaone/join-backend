package types

import (
	"github.com/google/uuid"
	"time"
)

type Code struct {
	// Invite code itself
	Code uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`

	// Who created this code
	CreatedByUserID uint `gorm:"index;column:created_by_user_id"`

	// Props
	CodeProps
}

type CodeProps struct {
	// Comment for this code
	Comment string `json:"comment"`

	// Status, could be deactivated and reactivated by creator
	IsActivate bool `json:"is_activate"`

	// Limitations
	RegisterCountLimit     int64     `json:"register_count_limit"` // 0 means no limit (!!not strict!!)
	RegisterTimeStart      time.Time `json:"register_time_start"`  // Start from now if not specified
	RegisterTimeEnd        time.Time `json:"register_time_end"`
	IsRegisterTimeEndValid bool      `json:"is_register_time_end_valid"` // Set to false if ends time is not limited
	RegisterCoolDown       uint      `json:"register_cool_down"`         // In seconds, 0 means no cooldown
}
