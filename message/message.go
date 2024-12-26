package message

import "fmt"

// The payload data types must implement this interface
type Message interface {
    // Check if the message its valid
    Validate() error
    // Return the message type name. eg: "video"
    GetType() MessageType
}

// Media extends the Message interface
type Media interface {
    Message
    GetID() string
    GetLink() string
}

type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Image struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Video struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Document struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`
}

type Audio struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

type Sticker struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

func mediaBaseValidation(m Media) error {
    if (m.GetLink() != "") && (m.GetID() != "") {
        return ErrorIdAndLink
    }
    if (m.GetLink() == "") && (m.GetID() == "") {
        return ErrorIdAndLink
    }
    return nil
}

func (m *Text) Validate() error { 
    if m.Body == "" {
        return fmt.Errorf("text body cannot be empty")
    }
    return nil
}

// implements Media interface
func (m *Image)     GetID() string {return m.ID}
func (m *Video)     GetID() string {return m.ID}
func (m *Audio)     GetID() string {return m.ID}
func (m *Sticker)   GetID() string {return m.ID}
func (m *Document)  GetID() string {return m.ID}

func (m *Image)     GetLink() string {return m.Link}
func (m *Video)     GetLink() string {return m.Link}
func (m *Audio)     GetLink() string {return m.Link}
func (m *Sticker)   GetLink() string {return m.Link}
func (m *Document)  GetLink() string {return m.Link}

// implements the Message interface
func (m *Image)     Validate() error { return mediaBaseValidation(m)}
func (m *Video)     Validate() error { return mediaBaseValidation(m)}
func (m *Audio)     Validate() error { return mediaBaseValidation(m)}
func (m *Sticker)   Validate() error { return mediaBaseValidation(m)}
func (m *Document)  Validate() error { return mediaBaseValidation(m)}

func (m *Image)     GetType() MessageType { return ImageType }
func (m *Video)     GetType() MessageType { return VideoType }
func (m *Audio)     GetType() MessageType { return AudioType }
func (m *Sticker)   GetType() MessageType { return StickerType }
func (m *Document)  GetType() MessageType { return DocumentType }
func (m *Text)      GetType() MessageType { return TextType }
