package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

	data Message
}

func NewPayload(to string, data Message) *Payload {
	return &Payload{
		messagingProduct: "whatsapp",
		recipientType:    "individual",
		data:             data,
		to:               to,
	}
}

// For debugging purposes
func (p *Payload) String() string {
    t, _ := json.MarshalIndent(p, "", "   ")
    return fmt.Sprintf("%s\n", string(t))
}

func (p *Payload) Validate() error {
	if p.data == nil {
		return NewErr(p.data, ErrEmptyPayload)
	}
	if p.to == "" {
		return NewErr(p.data, ErrNoRecipient)
	}
    return p.data.Validate()
}

func (p *Payload) MarshalJSON() ([]byte, error) {
	typ := p.data.GetType()

	// Create a map to hold the payload data
	dataMap := map[string]any{
		"messaging_product": p.messagingProduct,
		"recipient_type":    p.recipientType,
		"to":                p.to,
		"type":              typ,
		string(typ):         p.data,
	}

	return json.Marshal(dataMap)
}

func New(accessToken, numberId string) *Whatsapp {
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

func (w *Whatsapp) Send(to string, msg Message) (*Response, error) {
	payload := NewPayload(to, msg)
	if err := payload.Validate(); err != nil {
		return nil, err
	}
	req, err := w.createReq(payload)
	if err != nil {
		return nil, err
	}

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

func (w *Whatsapp) createReq(p *Payload) (*http.Request, error) {
    var buf bytes.Buffer
	   err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, w.url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.accessToken))
	return req, nil
}

func (w *Whatsapp) SendText(to string, msg string) (*Response, error) {
	return w.Send(to, NewTextMessage(msg))
}
