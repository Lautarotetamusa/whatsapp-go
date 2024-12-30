package whatsapp

import (
	"encoding/json"
	"errors"
	"fmt"
)

type InteractiveAction interface {
	GetInteractionType() InteractionType
	Validate() error
}

type Interactive struct {
	Type InteractionType `json:"type"`

	Action Action `json:"action"`

	// Optional for ProductType, required for the other types
	Body *Body `json:"body,omitempty"`

	// Required for ProductList, optional for the other types
	Header *Header `json:"header,omitempty"`

	// Optional
	Footer *Footer `json:"footer,omitempty"`
}

// Holds the data of any type of interaction action
type Action struct {
	msg InteractiveAction
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages/#header-object
type Header struct {
	data Message
}

// Optional for ProductType, required for the other types
type Body struct {
	// Max lenght: 1024 chars
	// Emojis and markdown allowed
	Text string `json:"text"`
}

type Footer struct {
	// Max lenght: 60 chars
	// Emojis and markdown allowed
	Text string `json:"text"`
}

type CallToAction struct {
	Name       string     `json:"name"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	DisplayText string `json:"display_text"`
	// URL to load in the device's default web browser when tapped by the WhatsApp user.
	URL string `json:"url"`
}

// Max 3 buttons per message
type Buttons []Button
type Button struct {
	// "reply"
	Type string `json:"type"`

	Reply Reply `json:"reply"`
}
type Reply struct {
	// Must to be unique in one message
	// Max length: 20 chars
	Title string `json:"title"`

	// Max length: 256 chars
	// Cannot start or end with a space
	ID string `json:"id"`
}

func (btn *Button) Validate() error {
	if btn.Reply.Title == "" {
		return NewErr(&Interactive{}, errors.New("button title cannot be empty"))
	}

	if len(btn.Reply.Title) > 20 {
		return NewErr(&Interactive{}, errors.New("button title cannot have more than 20 characters"))
	}

	if len(btn.Reply.ID) > 256 {
		return NewErr(&Interactive{}, errors.New("button ID cannot have more than 256 characters"))
	}

	if btn.Reply.ID[0] == ' ' || btn.Reply.ID[len(btn.Reply.ID)-1] == ' ' {
		return NewErr(&Interactive{}, errors.New("button ID cannot start or end with a space"))
	}

	return nil
}

func (buttons Buttons) Validate() error {
	if len(buttons) > 3 {
		return NewErr(&Interactive{}, errors.New("interaction cannot have more than 3 buttons"))
	}

	idMap := make(map[string]bool, 3)
	for _, btn := range buttons {
		btn.Validate()

		if idMap[btn.Reply.ID] {
			return NewErr(&Interactive{}, errors.New("two buttons cannot have the same ID in one message"))
		}
		idMap[btn.Reply.ID] = true
	}
	return nil
}

func (buttons Buttons) GetInteractionType() InteractionType {
	return ButtonsType
}

func NewButton(title, id string) Button {
	return Button{
		Type: "reply",
		Reply: Reply{
			Title: title,
			ID:    id,
		},
	}
}

func NewButtons(buttons ...Button) Buttons {
	return Buttons(buttons)
}

func (cta CallToAction) GetInteractionType() InteractionType {
	return CallToActionType
}

func (cta CallToAction) Validate() error {
	// TODO:
	return nil
}

func (h *Header) Validate() error {
	if h.data == nil {
		return NewErr(h.data, ErrEmptyPayload)
	}
	return nil
}

func (a *Action) MarshalJSON() ([]byte, error) {
	typ := a.msg.GetInteractionType()
	dataMap := map[string]any{
		"name":      typ,
		string(typ): a.msg,
	}

	return json.Marshal(dataMap)
}

func (h *Header) MarshalJSON() ([]byte, error) {
	typ := h.data.GetType()
	dataMap := map[string]any{
		"type":      typ,
		string(typ): h.data,
	}

	return json.Marshal(dataMap)
}

func (i *Interactive) Validate() error {
	typ := i.Action.msg.GetInteractionType()
	if typ != Product {
		if i.Body == nil {
			return NewErr(i, fmt.Errorf("body its required in '%s' type", typ))
		}
	}
	return i.Action.msg.Validate()
}

func (i *Interactive) GetType() MessageType {
	return InteractiveType
}

func NewInteractive(action InteractiveAction) *Interactive {
	return &Interactive{
		Type: ButtonType,
		Action: Action{
			msg: action,
		},
	}
}

func NewCallToAction(body, displayText, url string) *Interactive {
	return &Interactive{
		Action: Action{
			msg: CallToAction{
				Name: string(CallToActionType),
				Parameters: Parameters{
					DisplayText: displayText,
					URL:         url,
				},
			},
		},
		// This its not required in all types
		Body: &Body{
			Text: body,
		},
	}
}

func (i *Interactive) SetBody(text string) *Interactive {
	i.Body = &Body{
		Text: text,
	}
	return i
}

func (i *Interactive) SetHeader(msg Message) *Interactive {
	i.Header = &Header{
		data: msg,
	}
	return i
}

func (i *Interactive) SetFooter(text string) *Interactive {
	i.Footer = &Footer{
		Text: text,
	}
	return i
}

func NewBody(text string) *Body {
	return &Body{
		Text: text,
	}
}

func NewFooter(text string) *Footer {
	return &Footer{
		Text: text,
	}
}
