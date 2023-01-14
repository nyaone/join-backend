package misskey

import (
	"join-nyaone/config"
)

type AccountsCreate_Request struct {
	I        string `json:"i"` // Only works with ADMIN permission! Moderator's don't work.
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountsCreate_Response struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func CreateAccount(username string, password string) (*AccountsCreate_Response, error) {

	return PostAPIRequest[AccountsCreate_Response]("/api/admin/accounts/create", &AccountsCreate_Request{
		I:        config.Config.Misskey.Token.Admin,
		Username: username,
		Password: password,
	})

}
