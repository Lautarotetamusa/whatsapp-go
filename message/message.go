package message

import (
	"fmt"
	"strings"
)

// The payload data types must implement this interface
type Message interface {
    // Check if the message its valid
    Validate() error
    // Return the message type name. eg: "video"
    GetType() MessageType
}

// General Media type
type Media struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Image struct {
    *Media
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Video struct {
    *Media
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Document struct {
    *Media
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`
}

type Audio struct {
    *Media
}

type Sticker struct {
    *Media
}

func main(){
    m := Image{
        Caption: "",
        Media: FromID("123"),
    } 

    fmt.Println(m)
}

func isUrl(s string) bool {
	return strings.Contains(s, "http://") || strings.Contains(s, "https://")
}

func NewTextMessage(msg string) *Text {
	return &Text{
        // if the message its an url make the message shows the preview
		PreviewUrl: isUrl(msg),
		Body:       msg,
	}
}

func NewMedia(idOrLink string) *Media {
    m := Media{}
    if isUrl(idOrLink){
        m.Link = idOrLink
    }else{
        m.ID = idOrLink
    }
    return &m
}

func FromID(id string) *Media {
    return &Media{
        ID: id,
    }
}

func FromLink(link string) *Media {
    return &Media{
        Link: link,
    }
}

func (m *Media) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return ErrorIdAndLink
    }
    if (m.Link == "") && (m.ID == "") {
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

// implements the Message interface
func (m *Image)     GetType() MessageType { return ImageType }
func (m *Video)     GetType() MessageType { return VideoType }
func (m *Audio)     GetType() MessageType { return AudioType }
func (m *Sticker)   GetType() MessageType { return StickerType }
func (m *Document)  GetType() MessageType { return DocumentType }
func (m *Text)      GetType() MessageType { return TextType }
