package types

import "github.com/lib/pq"

type ApplicationPublic struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CallbackURL string         `json:"callbackUrl"`
	Permission  pq.StringArray `json:"permission" gorm:"type:text[]"`
}

type ApplicationPrivate struct {
	AppID  string `json:"id"`
	Secret string `json:"secret"`
}
