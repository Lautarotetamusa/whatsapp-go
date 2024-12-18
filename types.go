package whatsapp

import (
	"net/http"

	"log/slog"
)

type Whatsapp struct {
	accessToken string
	numberId    string
	client      *http.Client
	url         string
    logger      *slog.Logger
}

type Response struct {
	Contacts []struct {
		Input string `json:"input"`
		WaId  string `json:"wa_id"`
	}
	Messages []struct {
		Id string `json:"id"`
	}
}

type TextPayload struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type MediaPayload struct {
	Id string `json:"id"`
}

type DocumentPayload struct {
	Link     string `json:"link"`
	Caption  string `json:"caption"`
	Filename string `json:"filename"`
}

type TemplatePayload struct {
	Name       string       `json:"name"`
    Language   Language     `json:"language"`
	Components []Components `json:"components"`
}

type Parameter struct {
    Type        string           `json:"type"`

    Text        *TextPayload     `json:"text,omitempty"`
	Image       *MediaPayload    `json:"image,omitempty"`
	Video       *MediaPayload    `json:"video,omitempty"`
	Document    *DocumentPayload `json:"document,omitempty"`
}

type Components struct {
    Type       string       `json:"type"`
	Parameters []Parameter  `json:"parameters"`
}

type Language struct {
	Code string `json:"code"`
}

type Payload struct {
	MessagingProduct string           `default:"whatsapp" json:"messaging_product"`
	RecipientType    string           `default:"individual" json:"recipient_type"`
	To               string           `json:"to"`
	Type             MessageType      `json:"type"`

	Text             *TextPayload     `json:"text,omitempty"`
	Template         *TemplatePayload `json:"template,omitempty"`
	Image            *MediaPayload    `json:"image,omitempty"`
	Video            *MediaPayload    `json:"video,omitempty"`
	Document         *DocumentPayload `json:"document,omitempty"`
}
