package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"join-nyaone/config"
	"join-nyaone/global"
	"net/http"
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
	// Prepare request
	apiEndpoint := fmt.Sprintf("%s/api/notes/create", config.Config.Misskey.Instance)

	reqBodyBytes, err := json.Marshal(&NotesCreate_Request{
		I:    config.Config.Misskey.Token.Notify,
		Text: text,
	})
	if err != nil {
		global.Logger.Errorf("Failed to marshall request body with error: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(reqBodyBytes))
	if err != nil {
		global.Logger.Errorf("Failed to prepare request with error: %v", err)
		return nil, err
	}

	// Do request
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to finish request with error: %v", err)
		return nil, err
	}

	// Parse response
	if res.StatusCode == http.StatusOK {
		var resBody NotesCreate_Response
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			global.Logger.Errorf("Failed to decode response body with error: %v", err)
			return nil, err
		}

		return &resBody, nil
	} else {
		global.Logger.Errorf("Request failed.")
		var errBody Error_Response
		err = json.NewDecoder(res.Body).Decode(&errBody)
		if err != nil {
			global.Logger.Errorf("Failed to decode error body with error: %v", err)
			return nil, err
		}

		global.Logger.Errorf("Failed details: %v", errBody)
		return nil, fmt.Errorf(errBody.Error.Message)
	}

}
