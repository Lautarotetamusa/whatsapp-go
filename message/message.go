package message

import (
	"fmt"
)

// The payload data types must implement this interface
type Message interface {
    // Check if the message its valid
    Validate() error
    // Return the message type name. eg: "video"
    TypeName() MessageType
}

type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Image struct {
    // Optional
    Caption string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

type Video struct {
    // Optional
    Caption string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

type Document struct {
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`

    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

func (m *Image) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return ErrorIdAndLink
    }
    return nil
}

func (m *Video) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return ErrorIdAndLink
    }
    return nil
}

func (m *Document) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return ErrorIdAndLink
    }
    return nil
}

func (m *Text) Validate() error {
    if m.Body == "" {
        return fmt.Errorf("Text body cannot be empty")
    }
    return nil
}

func (m *Image) TypeName() MessageType {
    return ImageType
}

func (m *Video) TypeName() MessageType {
    return VideoType
}

func (m *Document) TypeName() MessageType {
    return DocumentType
}

func (m *Text) TypeName() MessageType {
    return TextType
}
