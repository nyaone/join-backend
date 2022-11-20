package types

import "github.com/google/uuid"

type User struct {
	Username string `gorm:"uniqueIndex"`

	// Invited by, 0 means created by login rather than invite (from this system)
	InvitedByCode   uuid.UUID `gorm:"index;type:uuid;column:invited_by_code"`
	InvitedByUserID uint      `gorm:"index;column:invited_by_user_id"`
}
