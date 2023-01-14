package misskey

import (
	"join-nyaone/config"
)

type NotesCreate_Request struct {
	I    string `json:"i"`
	Text string `json:"text"`
}

type NotesCreate_Response struct {
	CreatedNote struct {
		ID string `json:"id"`
	} `json:"createdNote"`
}

func NotesCreate(text string) (*NotesCreate_Response, error) {

	return PostAPIRequest[NotesCreate_Response]("/api/notes/create", &NotesCreate_Request{
		I:    config.Config.Misskey.Token.Notify,
		Text: text,
	})

}
