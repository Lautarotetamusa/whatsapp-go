package message

import (
	"fmt"
)

// The payload data types must implement this interface
type Message interface {
    // Check if the message its valid
    Validate() error
    // Return the message type name. eg: "video"
    TypeName() string
}

type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Image struct {
    // Optional
    Caption string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link"`
	ID      string     `json:"id"`
}

type Video struct {
    // Optional
    Caption string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link"`
	ID      string     `json:"id"`
}

type Document struct {
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link"`
	ID      string     `json:"id"`
}

func (m *Image) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return fmt.Errorf(ErrorIdAndLink, m.TypeName())
    }
    return nil
}

func (m *Video) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return fmt.Errorf(ErrorIdAndLink, m.TypeName())
    }
    return nil
}

func (m *Document) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return fmt.Errorf(ErrorIdAndLink, m.TypeName())
    }
    return nil
}

func (m *Text) Validate() error {
    if m.Body == "" {
        return fmt.Errorf("Text body cannot be empty")
    }
    return nil
}

func (m *Image) TypeName() string {
    return ImageType
}

func (m *Video) TypeName() string {
    return VideoType
}

func (m *Document) TypeName() string {
    return DocumentType
}

func (m *Text) TypeName() string {
    return TextType
}
