package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager
type InputValidation interface {
	Validate() error
}

var (
	// ErrInvalidName - error when we have an invalid name
	UserNameReq = errors.New("message empty or invalid")
)

func (t MessagesCreateInput) Validate() error {
	if govalidator.IsNull(t.Message) {
		return UserNameReq
	}
	return nil
}

