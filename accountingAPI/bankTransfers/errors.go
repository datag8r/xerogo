package banktransfers

import "errors"

var (
	ErrInvalidBankTransferID          = errors.New("invalid bank transfer id")
	ErrInvalidBankTransferForCreation = errors.New("invalid bank transfer id")
)
