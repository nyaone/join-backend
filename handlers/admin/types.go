package admin

import (
	"join-nyaone/types"
	"time"
)

type CodeResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	types.CodeProps

	InviteCount int64 `json:"invite_count"`
}

type InviteesResponse struct {
	RegisteredAt time.Time `json:"registered_at"`

	Username        string `json:"username"`
	InvitedByCodeID uint   `json:"invited_by_code"`
}
