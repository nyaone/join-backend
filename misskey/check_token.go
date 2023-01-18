package misskey

import (
	"time"
)

type I_Request struct {
	I string `json:"i"`
}

type MisskeyUser struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	AvatarUrl   string    `json:"avatarUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	IsAdmin     bool      `json:"isAdmin"`
	IsModerator bool      `json:"isModerator"`
	// Ignore other fields
}

type I_Response MisskeyUser

func CheckToken(token string) (*I_Response, error) {

	return PostAPIRequest[I_Response]("/api/i", &I_Request{
		I: token,
	})

}
