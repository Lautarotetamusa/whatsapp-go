package message

type MessageType string
const (
    TextType        MessageType = "text"
    ImageType       MessageType = "image"
    AudioType       MessageType = "audio"
    StickerType     MessageType = "sticker"
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
