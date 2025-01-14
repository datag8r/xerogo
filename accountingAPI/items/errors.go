package items

import "errors"

var (
	ErrInvalidItemForCreation = errors.New("some fields are invalid to create an item")
	ErrInvalidItemForUpdating = errors.New("some fields are invalid to update an item" +
		" - usually this is the code field")
	ErrInvalidItemID = errors.New("invalid item id for request")
)
