package whatsapp

import (
	"errors"
	"fmt"
)

type ResponseError struct {
	Code      int    `json:"code"`
	FbtraceId string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

type ValidationError struct {
    msg Message
    err error
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("(%d) %s - %s\nfbtrace_id: %s", e.Code, e.Type, e.Message, e.FbtraceId)
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

