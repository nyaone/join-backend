package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"join-nyaone/config"
	"join-nyaone/global"
	"net/http"
)

type MessagingMessageCreate_Request struct {
	I      string `json:"i"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type MessagingMessageCreate_Response struct {
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}

func SendMessage(targetUserID string, text string) (*MessagingMessageCreate_Response, error) {
	// Prepare request
	apiEndpoint := fmt.Sprintf("%s/api/messaging/messages/create", config.Config.Misskey.Instance)

	reqBodyBytes, err := json.Marshal(&MessagingMessageCreate_Request{
		I:      config.Config.Misskey.Token.Notify,
		UserID: targetUserID,
		Text:   text,
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
		var resBody MessagingMessageCreate_Response
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
