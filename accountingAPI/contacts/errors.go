package contacts

import "errors"

var (
	ErrInvalidContactForCreation = errors.New("one or more required fields were invalid to create this contact" +
		" - name field required")
	ErrInvalidContactForUpdating = errors.New("a valid contact id is required to update an contact")
	ErrInvalidContactID          = errors.New("invalid contact id for request")
)
