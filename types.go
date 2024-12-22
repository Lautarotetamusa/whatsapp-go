package whatsapp

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Whatsapp struct {
	accessToken string
	numberId    string
	client      *http.Client
	url         string
}

type Response struct {
	Contacts []struct {
		Input string `json:"input"`
		WaId  string `json:"wa_id"`
	}
	Messages []struct {
		Id string `json:"id"`
	}
    Error *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
    Code        int `json:"code"`
    FbtraceId   string  `json:"fbtrace_id"`
    Message     string  `json:"message"`
    Type        string  `json:"type"`
}

type TextPayload struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type MediaPayload struct {
    // Optional
    Caption string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link"`
	ID      string     `json:"id"`
}

type DocumentPayload struct {
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link"`
	ID      string     `json:"id"`
}

// The payload data types must implement this interface
type PayloadType interface {
    validate() error
}

type Payload struct {
	MessagingProduct string           `default:"whatsapp" json:"messaging_product"`
	RecipientType    string           `default:"individual" json:"recipient_type"`
	To               string           `json:"to"`
	Type             MessageType      `json:"type"`

    Data             PayloadType
}

func (p *Payload) MarshalJSON() ([]byte, error) {
    // Create a map to hold the payload data
    dataMap := map[string]any{
        "messaging_product": p.MessagingProduct,
        "recipient_type":    p.RecipientType,
        "to":                p.To,
        "type":              p.Type.String(),
    }

    // Add the Data field to the map
    if p.Data != nil {
        dataValue := reflect.ValueOf(p.Data)
        if dataValue.Kind() == reflect.Ptr {
            dataValue = dataValue.Elem()
        }
        dataMap[p.Type.String()] = dataValue.Interface()
    }

    return json.Marshal(dataMap)
}
