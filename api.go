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

type Payload struct {
	messagingProduct string 
	recipientType    string
	to               string

    data             message.Message
}

func NewPayload(data message.Message) *Payload {
	return &Payload{
        messagingProduct: "whatsapp",
        recipientType: "individual",
        data: data,
	}
}

func (p *Payload) MarshalJSON() ([]byte, error) {
    if p.data == nil {
        return nil, fmt.Errorf("payload data its empty")
    }
    typeName := p.data.TypeName() 

    // Create a map to hold the payload data
    dataMap := map[string]any{
        "messaging_product": p.messagingProduct,
        "recipient_type":    p.recipientType,
        "to":                p.to,
        "type":              typeName,
    }

    // Add the Data field to the map
    dataValue := reflect.ValueOf(p.data)
    if dataValue.Kind() == reflect.Ptr {
        dataValue = dataValue.Elem()
    }
    dataMap[string(typeName)] = dataValue.Interface()

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

func NewWhatsapp(accessToken, numberId string) *Whatsapp {
    if accessToken == "" || numberId == "" {
        panic("accessToken and numberId cannot be empty")
    }
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
    if to == "" {
        return nil, fmt.Errorf("recipient phone cannot be empty")
    }
    payload := NewPayload(msg)
    payload.to = to

	jsonBody, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    if err := payload.data.Validate(); err != nil {
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
        return nil, data.Error
    }

    if res.StatusCode != http.StatusOK {
        return nil, data.Error
    }

	return &data, nil
}

func (w *Whatsapp) SendText(to string, message string) (*Response, error) {
	return w.Send(to, newTextPayload(message))
}
