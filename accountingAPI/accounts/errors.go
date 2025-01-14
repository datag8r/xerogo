package accounts

import "errors"

var (
	ErrInvalidAccountForCreation = errors.New("one or more required fields were invalid to create this account")
	ErrInvalidAccountForUpdating = errors.New("a valid account id is required to update an account")
	ErrInvalidAccountID          = errors.New("invalid account id for request")
)
