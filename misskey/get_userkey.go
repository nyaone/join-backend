package misskey

type AuthSessionUserkey_Request struct {
	AppSecret string `json:"appSecret"`
	Token     string `json:"token"`
}

type AuthSessionUserkey_Response struct {
	AccessToken string      `json:"accessToken"`
	User        MisskeyUser `json:"user"`
}

func GetUserkey(appSecret string, token string) (*AuthSessionUserkey_Response, error) {

	return PostAPIRequest[AuthSessionUserkey_Response]("/api/auth/session/userkey", &AuthSessionUserkey_Request{
		AppSecret: appSecret,
		Token:     token,
	})

}
