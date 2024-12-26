package whatsapp

import "fmt"

type ResponseError struct {
	Code      int    `json:"code"`
	FbtraceId string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("(%d) %s - %s\nfbtrace_id: %s", e.Code, e.Type, e.Message, e.FbtraceId)
}
