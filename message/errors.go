package message

import (
	"errors"
	"fmt"
)

type ValidationError struct {
    msg Message
    err error
}

func NewErr(m Message, e error) *ValidationError {
    return &ValidationError{
        msg: m,
        err: e,
    }
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("type: %s err: %s", e.msg.GetType(), e.err.Error())
}

var (
    ErrorIdAndLink = errors.New("either Link or ID must be present in media message")
)

