package errors

import (
	"errors"
)

var (
	ErrEndpointCallNotSupported = errors.New("this endpoint does not support this call type")
	ErrNoPageDataReturned       = errors.New("no page data returned")
)
