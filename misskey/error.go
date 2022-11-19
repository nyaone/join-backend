package misskey

type Error_Response struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		ID      string `json:"id"`
		Kind    string `json:"kind"`
	} `json:"error"`
}
