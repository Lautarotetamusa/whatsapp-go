package message

type PayloadError string

func (e PayloadError) Error() string {
    return string(e)
}

const (
    ErrorIdAndLink PayloadError = "either Link or ID must be present in media message"
)

