package message

import (
	"fmt"
)

const (
    ErrorIdAndLink string = "either Link or ID must be present in %s message"
)

type messageNotValidError struct {
    mType   string
    message string
}

func (e *messageNotValidError) Error() string {
    return fmt.Sprintf("%s - %s", e.mType, e.message)
}

