package misskey

import (
	"fmt"
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

	result, err := PostAPIRequest[UsernameAvailable_Response]("/api/username/available", &UsernameAvailable_Request{
		Username: username,
	})
	if err != nil {
		return false, err
	} else {
		return result.Available, nil
	}
}
