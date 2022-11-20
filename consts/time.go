package consts

import "time"

const (
	TIME_LOGIN_REQUEST_VALID = 3 * time.Minute
	TIME_LOGIN_SESSION_VALID = 6 * time.Hour

	TIME_NEW_USER_SEND_INVITATION_AFTER = 24 * time.Hour
)
