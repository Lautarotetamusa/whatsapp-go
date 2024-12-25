package message

import (
	"encoding/json"
	"fmt"
)

type MessageType string
const (
    TextType        MessageType = "text"
    ImageType       MessageType = "image"
    VideoType       MessageType = "video"
    DocumentType    MessageType = "document"
    TemplateType    MessageType = "template"
)

// "header" or "body"
type ComponentType string
const (
    ComponentTypeHeader ComponentType = "header"
    ComponentTypeBody   ComponentType = "body"
)

func (code LanguageCode) MarshalJSON() ([]byte, error) {
    if _, ok := langCodeValue[string(code)]; !ok {
        return nil, fmt.Errorf("%s isnt a valid Language code", code)
    }
    return json.Marshal(code)
}
