package whatsapp

import (
	"encoding/json"
	"errors"
)

// Holds the data of any type of interaction action
// Any type of action of an Interactive message 
// Each action have a specific required field
type Action interface {
	GetInteractionType() InteractionType

	Validate() error
}

type Interactive struct {
	action Action

	// Optional for ProductType, required for the other types
	Body *Body `json:"body,omitempty"`

	// Required for ProductList, optional for the other types
	Header *Header `json:"header,omitempty"`

	// Optional
	Footer *Footer `json:"footer,omitempty"`
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages/#header-object
type Header struct {
	data Message
}

// Optional for ProductType, required for the other types
type Body struct {
	// Max length: 1024 chars
	// Emojis and markdown allowed
	Text string `json:"text"`
}

type Footer struct {
	// Max length: 60 chars
	// Emojis and markdown allowed
	Text string `json:"text"`
}

type ListSection struct {
    // Required if the message have more than one section
    Title   string  `json:"title"`
    // Max len: 10
    Rows []Row `json:"rows"`
}

type ProductSection struct {
    // Required if the message have more than one section
    Title   string  `json:"title"`
    // Max len: 30
    ProductItems []Product `json:"product_items"`
}

type Product struct {
    ProductRetailerID string `json:"product_retailer_id"`
}

type Row struct {
    // Max length: 200 chars
    ID          string  `json:"id"`
    // Max length: 24 chars
    Title       string  `json:"title"`
    // Optional. Max length: 72 chars
    Description string  `json:"description,omitempty"`
}

type CallToAction struct {
    Name        string      `json:"name"`
    Parameters  Parameters  `json:"parameters"`
}

type Parameters struct {
	DisplayText string `json:"display_text"`
	// URL to load in the device's default web browser when tapped by the WhatsApp user.
	URL string `json:"url"`
}

// Max 3 buttons per message
type Buttons struct {
    Buttons []Button `json:"buttons"`
}
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

type List struct {
    // Max length: 20 chars
    Button      string        `json:"button"`
    Sections    []ListSection `json:"sections"`
}

func (r *Row) Validate() error {
    if len(r.Title) > 24 {
        return NewErr(r, errors.New("row title cannot have more than 24 characters"))
    }

    if len(r.ID) > 200 {
        return NewErr(r, errors.New("row id cannot have more than 200 characters"))
    }

    if r.Description != "" && len(r.Description) > 72 {
        return NewErr(r, errors.New("row description cannot have more than 72 characters"))
    }

    return nil
}

func (s ListSection) Validate() error {
    if len(s.Rows) > 10 {
        return NewErr(s, errors.New("list section cannot have more than 10 rows"))
    }
    for _, row := range s.Rows {
        if err := row.Validate(); err != nil {
            return err
        }
    }
    return nil
}

func (l List) Validate() error {
    if len(l.Button) > 20 {
        return NewErr(l, errors.New("list button cannot have more than 20 chars"))
    }
    if len(l.Sections) == 0 {
        return NewErr(l, errors.New("list must have at least 1 section"))
    }
    if len(l.Sections) > 10 {
        return NewErr(l, errors.New("list cannot have more than 10 sections"))
    }
    for _, section := range l.Sections {
        if err := section.Validate(); err != nil {
            return err
        }
    }
    return nil
}

func (l List) GetInteractionType() InteractionType {
    return ListType
}

func NewList(button string, listSections ...ListSection) List {
    return List{
        Button: button,
        Sections: listSections,
    }
}

func NewListSection(title string, rows ...Row) ListSection {
    return ListSection{
        Title: title,
        Rows: rows,
    }
}

func NewRow(id, title, desc string) Row {
    return Row{
        ID: id,
        Title: title,
        Description: desc,
    }
}

func (btn *Button) Validate() error {
	if btn.Reply.Title == "" {
		return NewErr(btn, errors.New("button title cannot be empty"))
	}

	if len(btn.Reply.Title) > 20 {
		return NewErr(btn, errors.New("button title cannot have more than 20 characters"))
	}

	if len(btn.Reply.ID) > 256 {
		return NewErr(btn, errors.New("button ID cannot have more than 256 characters"))
	}

	if btn.Reply.ID[0] == ' ' || btn.Reply.ID[len(btn.Reply.ID)-1] == ' ' {
		return NewErr(btn, errors.New("button ID cannot start or end with a space"))
	}

	return nil
}

func (buttons Buttons) Validate() error {
	if len(buttons.Buttons) > 3 {
		return NewErr(buttons, errors.New("interaction cannot have more than 3 buttons"))
	}

	idMap := make(map[string]bool, 3)
	for _, btn := range buttons.Buttons {
		if err := btn.Validate(); err != nil {
			return err
		}

		if idMap[btn.Reply.ID] {
			return NewErr(buttons, errors.New("two buttons cannot have the same ID in one message"))
		}
		idMap[btn.Reply.ID] = true
	}
	return nil
}

func (buttons Buttons) GetInteractionType() InteractionType {
	return ButtonType
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
	return Buttons{
        Buttons: buttons,
    }
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

// TODO: Maybe can do this better??
func (i *Interactive) MarshalJSON() ([]byte, error) {
	typ := i.action.GetInteractionType()
	dataMap := map[string]any{
        "type": typ,
        "action": i.action,
	}
    if i.Body != nil {
        dataMap["body"] = i.Body
    }
    if i.Header != nil {
        dataMap["header"] = i.Header
    }
    if i.Footer != nil {
        dataMap["footer"] = i.Footer
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
	typ := i.action.GetInteractionType()
	if typ != ProductType {
		if i.Body == nil {
			return NewErr(i.action, errors.New("body its required"))
		}
	}
	return i.action.Validate()
}

func (i *Interactive) GetType() MessageType {
	return InteractiveType
}

func NewInteractive(action Action) *Interactive {
	return &Interactive{
        action: action,
    }
}

func NewCallToAction(body, displayText, url string) *Interactive {
	return &Interactive{
		action: CallToAction{
            Name: "cta_url",
            Parameters: Parameters{
                DisplayText: displayText,
                URL:         url,
            },
		},
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
