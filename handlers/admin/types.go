package admin

import (
	"join-nyaone/types"
	"time"
)

type CodeResponse struct {
	Code string `json:"code"`
	types.CodeProps

	InviteCount int64 `json:"invite_count"`
	IsValid     bool  `json:"is_valid"`
}

type InviteesResponse struct {
	RegisteredAt time.Time `json:"registered_at"`

	Username      string `json:"username"`
	InvitedByCode string `json:"invited_by_code"`
}
