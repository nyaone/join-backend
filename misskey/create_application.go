package misskey

import (
	"join-nyaone/types"
)

type AppCreate_Request types.ApplicationPublic

type AppCreate_Response struct {
	types.ApplicationPublic
	types.ApplicationPrivate
}

func CreateApplication(appMeta *types.ApplicationPublic) (*AppCreate_Response, error) {

	return PostAPIRequest[AppCreate_Response]("/api/app/create", appMeta)

}
