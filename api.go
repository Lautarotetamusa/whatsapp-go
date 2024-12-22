package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const version = "v21.0"
const baseUrl = "https://graph.facebook.com/%s/%s/messages"

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

func newPayload(t MessageType) *Payload {
	return &Payload{
        MessagingProduct: "whatsapp",
        RecipientType: "individual",
		Type: t,
	}
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
func newTextPayload(message string) *Payload {
	p := newPayload(TextMessage)

    previewUrl := strings.Contains(message, "http://") || strings.Contains(message, "https://")

	p.Data = &TextPayload{
		PreviewUrl: previewUrl,
		Body:       message,
	}
	return p
}

func NewDocumentPayload(link string, caption string, filename string) *Payload {
	p := newPayload(DocumentMessage)
    doc := DocumentPayload{
		Link:     link,
		Caption:  caption,
		Filename: filename,
	}
    p.Data = &doc
	return p
}

func (m *MediaPayload) validate() error {
    if (m.Link != "") && (m.ID != "") {
        return fmt.Errorf(ErrorIdAndLink, ImageMessage.String())
    }
    return nil
}

func (m *DocumentPayload) validate() error {
    if (m.Link != "") && (m.ID != "") {
        return fmt.Errorf(ErrorIdAndLink, DocumentMessage.String())
    }
    return nil
}
func (m *TextPayload) validate() error {
    return nil
}

func (w *Whatsapp) Send(to string, payload *Payload) (*Response, error) {
    payload.To = to
	jsonBody, _ := json.Marshal(payload)

    if err := payload.Data.validate(); err != nil {
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

func (w *Whatsapp) SendImage(to string, media MediaPayload) (*Response, error){
    p := newPayload(ImageMessage)
    p.Data = &media
	return w.Send(to, p)
}

func (w *Whatsapp) SendVideo(to string, media MediaPayload) (*Response, error) {
    p := newPayload(VideoMessage)
    p.Data = &media
	return w.Send(to, p)
}

func (w *Whatsapp) SendDocument(to string, doc DocumentPayload) (*Response, error) {
    p := newPayload(DocumentMessage)
    p.Data = &doc
	return w.Send(to, p)
}
