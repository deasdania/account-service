package utilities

import "os"

const (
	// endpoints
	CREATE_ACCOUNT  = "/create/account"
	TEST            = "/test"
	GET_ALL_ACCOUNT = "/users"
	GET_ACCOUNT     = "/user"
	LOGIN           = "/login"

	GENERATE_UUID = "/generate/uuid"
)

var (
	KEY_JWT = os.Getenv("SECRET_KEY_JWT")
)
