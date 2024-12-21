package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const version = "v21.0"
const baseUrl = "https://graph.facebook.com/%s/%s/messages"

func NewWhatsapp(accessToken string, numberId string, l *slog.Logger) *Whatsapp {
	return &Whatsapp{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		accessToken: accessToken,
		numberId:    numberId,
		url:         fmt.Sprintf(baseUrl, version, numberId),
        logger:      l.With("module", "whatsapp"),
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

	p.Text = &TextPayload{
		PreviewUrl: previewUrl,
		Body:       message,
	}
	return p
}

func NewTemplatePayload(name string, components []Components) *Payload {
	p := newPayload(TextMessage)
	p.Template = &TemplatePayload{
		Name:       name,
		Components: components,
		Language: Language{
			Code: "es_MX",
		},
	}
	return p
}

func NewMediaPayload(id string, t MessageType) *Payload {
    //TODO: 
	if t != ImageMessage && t != VideoMessage {
		panic("Type must be 'video' or 'image")
	}
	p := newPayload(t)

	m := &MediaPayload{
		Id: id,
	}
	if t == ImageMessage {
		p.Image = m
	} else {
		p.Video = m
	}
	return p
}

func NewDocumentPayload(link string, caption string, filename string) *Payload {
	p := newPayload(DocumentMessage)
	p.Document = &DocumentPayload{
		Link:     link,
		Caption:  caption,
		Filename: filename,
	}
	return p
}

func (w *Whatsapp) Send(to string, payload *Payload) (*Response, error) {
    payload.To = to

	jsonBody, _ := json.Marshal(payload)
    fmt.Println(string(jsonBody))
	w.logger.Debug("Sending message", "to", payload.To, "t", payload.Type)
	bodyReader := bytes.NewReader(jsonBody)

    fmt.Println(w.url)
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

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("message not sended status code: %s", res.Status)
    }

	var data Response
    if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
        return nil, err
    }

    // w.logger.Info("Message sended succesfully", "to", payload.To, "t", payload.Type, "id", data.Messages[0].Id)
	return &data, nil
}

func (w *Whatsapp) SendTemplate(to string, template TemplatePayload) (*Response, error) {
	p := newPayload(TemplateMessage)
	p.Template = &template
	return w.Send(to, p)
}

func (w *Whatsapp) SendText(to string, message string) (*Response, error) {
	return w.Send(to, newTextPayload(message))
}

func (w *Whatsapp) SendImage(to string, imageId string) {
	w.Send(to, NewMediaPayload(imageId, ImageMessage))
}

func (w *Whatsapp) SendVideo(to string, videoId string) {
	w.Send(to, NewMediaPayload(videoId, VideoMessage))
}
