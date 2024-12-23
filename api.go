package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/Lautarotetamusa/whatsapp-go/message"
)

const (
    version = "v21.0"
    baseUrl = "https://graph.facebook.com/%s/%s/messages"
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

type Payload struct {
	MessagingProduct string           `default:"whatsapp" json:"messaging_product"`
	RecipientType    string           `default:"individual" json:"recipient_type"`
	To               string           `json:"to"`

    Data             message.Message
}

func newPayload(data message.Message) *Payload {
	return &Payload{
        MessagingProduct: "whatsapp",
        RecipientType: "individual",
        Data: data,
	}
}

func (p *Payload) MarshalJSON() ([]byte, error) {
    if p.Data == nil {
        return nil, fmt.Errorf("payload data its empty")
    }
    typeName := p.Data.TypeName() 

    // Create a map to hold the payload data
    dataMap := map[string]any{
        "messaging_product": p.MessagingProduct,
        "recipient_type":    p.RecipientType,
        "to":                p.To,
        "type":              typeName,
    }

    // Add the Data field to the map
    dataValue := reflect.ValueOf(p.Data)
    if dataValue.Kind() == reflect.Ptr {
        dataValue = dataValue.Elem()
    }
    dataMap[typeName] = dataValue.Interface()

    return json.Marshal(dataMap)
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
func newTextPayload(msg string) *message.Text {
    previewUrl := strings.Contains(msg, "http://") || strings.Contains(msg, "https://")
	return &message.Text{
		PreviewUrl: previewUrl,
		Body:       msg,
    }
}

func NewWhatsapp(accessToken string, numberId string) *Whatsapp {
	return &Whatsapp{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		accessToken: accessToken,
		numberId:    numberId,
		url:         fmt.Sprintf(baseUrl, version, numberId),
	}
}

func (w *Whatsapp) Send(to string, msg message.Message) (*Response, error) {
    payload := newPayload(msg)
    payload.To = to

	jsonBody, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    if err := payload.Data.Validate(); err != nil {
        return nil, err
    }

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, w.url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.accessToken))

	res, err := w.client.Do(req)
	if err != nil {
        return nil, fmt.Errorf("Request error: %s", err)
	}
	defer res.Body.Close()

	var data Response
    if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
        return nil, err
    }

    if data.Error != nil {
        fmt.Println(data.Error)
        return nil, fmt.Errorf("%s %s", data.Error.Message, data.Error.Type)
    }

    if res.StatusCode != http.StatusOK {
        return &data, fmt.Errorf("message not sended status code: %s", res.Status)
    }

	return &data, nil
}

func (w *Whatsapp) SendText(to string, message string) (*Response, error) {
	return w.Send(to, newTextPayload(message))
}
