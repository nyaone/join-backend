package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"join-nyaone/config"
	"join-nyaone/global"
	"net/http"
	"time"
)

type UserShow_Request struct {
	Username string  `json:"username"`
	Host     *string `json:"host"` // Null
}

type UserShow_Response struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"` // Register time
	// Ignore others
}

func GetUser(username string) (*UserShow_Response, error) {
	// Check format
	if !usernameRegex.MatchString(username) {
		return nil, fmt.Errorf("invalid username format")
	}

	// Prepare request
	apiEndpoint := fmt.Sprintf("%s/api/users/show", config.Config.Misskey.Instance)

	reqBodyBytes, err := json.Marshal(&UserShow_Request{
		Username: username,
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
		var resBody UserShow_Response
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
