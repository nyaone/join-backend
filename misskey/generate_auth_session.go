package misskey

type AuthSessionGenerate_Request struct {
	AppSecret string `json:"appSecret"`
}

type AuthSessionGenerate_Response struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

func GenerateAuthSession(appSecret string) (*AuthSessionGenerate_Response, error) {

	return PostAPIRequest[AuthSessionGenerate_Response]("/api/auth/session/generate", &AuthSessionGenerate_Request{
		AppSecret: appSecret,
	})

}
