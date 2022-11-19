package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"join-nyaone/config"
	"join-nyaone/global"
	"net/http"
	"regexp"
)

type UsernameAvailable_Request struct {
	Username string `json:"username"`
}

type UsernameAvailable_Response struct {
	Available bool `json:"available"`
}

var (
	usernameRegex *regexp.Regexp
)

func init() {
	usernameRegex = regexp.MustCompile("^\\w{1,20}$")
}

func CheckUsername(username string) (bool, error) {
	// Check format
	if !usernameRegex.MatchString(username) {
		return false, fmt.Errorf("invalid username format")
	}

	// Prepare request
	apiEndpoint := fmt.Sprintf("%s/api/username/available", config.Config.Misskey.Instance)

	reqBodyBytes, err := json.Marshal(&UsernameAvailable_Request{
		Username: username,
	})
	if err != nil {
		global.Logger.Errorf("Failed to marshall request body with error: %v", err)
		return false, err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(reqBodyBytes))
	if err != nil {
		global.Logger.Errorf("Failed to prepare request with error: %v", err)
		return false, err
	}

	// Do request
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to finish request with error: %v", err)
		return false, err
	}

	// Parse response
	if res.StatusCode == http.StatusOK {
		var resBody UsernameAvailable_Response
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			global.Logger.Errorf("Failed to decode response body with error: %v", err)
			return false, err
		}

		return resBody.Available, nil
	} else {
		global.Logger.Errorf("Request failed.")
		var errBody Error_Response
		err = json.NewDecoder(res.Body).Decode(&errBody)
		if err != nil {
			global.Logger.Errorf("Failed to decode error body with error: %v", err)
			return false, err
		}

		global.Logger.Errorf("Failed details: %v", errBody)
		return false, fmt.Errorf(errBody.Error.Message)
	}
}
