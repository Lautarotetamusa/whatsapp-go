package whatsapp

import (
    "encoding/json"
    "fmt"
)

type MessageType int

// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components/#messages-object
const (
    TextMessage MessageType = iota
    AudioMessage
    ButtonMessage
    DocumentMessage
    ImageMessage
    InteractiveMessage
    OrderMessage
    StickerMessage
    VideoMessage
    SystemMessage
    UnknownMessageType
)

var (
    messageTypeName = map[MessageType]string{
        TextMessage: "text",
        AudioMessage: "audio",
        ButtonMessage: "button",
        DocumentMessage: "document",
        ImageMessage: "image",
        InteractiveMessage: "interactive",
        OrderMessage: "order",
        StickerMessage: "sticker",
        VideoMessage: "video",
        SystemMessage: "system",
        UnknownMessageType: "unknown",
    }
    messageTypeValue = map[string]MessageType{
        "text": TextMessage,
        "audio": AudioMessage,
        "button": ButtonMessage,
        "document": DocumentMessage,
        "image": ImageMessage,
        "interactive": InteractiveMessage,
        "order": OrderMessage,
        "sticker": StickerMessage,
        "video": VideoMessage,
        "system": SystemMessage,
        "unknown": UnknownMessageType,
    }
)

func ParseMessageType(s string) (MessageType, error) {
    value, ok := messageTypeValue[s]
    if !ok {
        return UnknownMessageType, fmt.Errorf("%q is not a valid message type", s)
    }
    return value, nil
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
