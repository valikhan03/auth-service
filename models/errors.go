package models

import(
	"errors"
)

var USER_ALREADY_EXISTS_ERR = errors.New("User with this email already exists")

var INVALID_USER_CREDS_ERR = errors.New("User's email or password is invalid")

var ERR_500 = errors.New("Internal server error (status code: 500)")