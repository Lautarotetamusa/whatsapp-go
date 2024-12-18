package whatsapp

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
    TextMessage MessageType = iota
    TemplateMessage
    ImageMessage
    VideoMessage
    DocumentMessage
    InteractiveMessage
    ButtonMessage
    OrderMessage
)

var (
    messageTypeName = map[MessageType]string{
        TextMessage: "text",
        TemplateMessage: "template",
        ImageMessage: "image",
        VideoMessage: "video",
        DocumentMessage: "document",
        InteractiveMessage: "interactive",
        ButtonMessage: "button",
        OrderMessage: "order", //TODO
    }
    messageTypeValue = map[string]MessageType{
        "text" : TextMessage,
        "template": TextMessage,
        "image": ImageMessage,
        "video": VideoMessage,
        "document": DocumentMessage,
        "interactive": InteractiveMessage,
        "button": ButtonMessage,
        "order": OrderMessage,
    }
)

func ParseMessageType(s string) (MessageType, error) {
	value, ok := messageTypeValue[s]
	if !ok {
		return MessageType(0), fmt.Errorf("%q is not a valid message type", s)
	}
	return MessageType(value), nil
}

func (s MessageType) String() string {
	return messageTypeName[s]
}

func (s *MessageType) UnmarshalJSON(data []byte) (err error) {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	if *s, err = ParseMessageType(status); err != nil {
		return err
	}
	return nil
}

func (s MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
