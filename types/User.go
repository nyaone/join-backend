package types

type User struct {
	Username string `gorm:"uniqueIndex"`

	// Invited by, 0 means created by login rather than invite (from this system)
	InvitedByCodeID uint `gorm:"index;column:invited_by_code_id"`
	InvitedByUserID uint `gorm:"index;column:invited_by_user_id"`
}
