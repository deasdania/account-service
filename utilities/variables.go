package utilities

import "os"

const (
	// endpoints
	CREATE_ACCOUNT  = "/create/account"
	TEST            = "/test"
	GET_ALL_ACCOUNT = "/users"
	GET_ACCOUNT     = "/user"
	// LOGIN           = "/login"
	CHANGE_PASSWORD = "/change/account/password"

	GENERATE_UUID = "/generate/uuid"

	LOGIN   = "/login"
	LOGOUT  = "/logout"
	REFRESH = "/refresh"
)

var (
	KEY_JWT      = os.Getenv("SECRET_KEY_JWT")
	ACCOUNT_PORT = os.Getenv("ACCOUNT_PORT")
	REDIS_URL    = os.Getenv("REDIS_URL")
)
