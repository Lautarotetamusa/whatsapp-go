package whatsapp

import (
	"errors"
	"fmt"
	"reflect"
)

type ResponseError struct {
	Code      int    `json:"code"`
	FbtraceId string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

type ValidationError struct {
	obj any
	err error
    typ reflect.Type
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("(%d) %s - %s\nfbtrace_id: %s", e.Code, e.Type, e.Message, e.FbtraceId)
}

func NewErr(m any, e error) *ValidationError {
	return &ValidationError{
		obj: m,
		err: e,
        typ: reflect.TypeOf(m),
	}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("type: %s err: %s", e.typ, e.err.Error())
}

var (
	ErrorIdAndLink  = errors.New("either Link or ID must be present in media message")
	ErrEmptyPayload = errors.New("payload data its empty")
	ErrNoRecipient  = errors.New("recipient phone cannot be empty")
)
