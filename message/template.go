package message

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://developers.facebook.com/docs/whatsapp/cloud-api/guides/send-message-templates/
type Template struct {
	Name       string       `json:"name"`
    Language   Language     `json:"language"`
	Components []Component  `json:"components"`
}

func (t *Template) Validate() error {
    // TODO
    return nil
}

func (t *Template) TypeName() MessageType {
    return TemplateType
}

type Parameter struct {
    // Type        MessageType         `json:"type"`
    //
    data        Message
    // Text        string           `json:"text,omitempty"`
	// Image       *MediaPayload    `json:"image,omitempty"`
	// Video       *MediaPayload    `json:"video,omitempty"`
	// Document    *DocumentPayload `json:"document,omitempty"`
}

type Component struct {
    Type       ComponentType    `json:"type"`

	Parameters []Parameter      `json:"parameters"`
}

type Language struct {
	Code string `json:"code"`
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
    if p.data == nil {
        return nil, fmt.Errorf("payload data its empty")
    }
    typeName := p.data.TypeName() 

    // Create a map to hold the payload data
    dataMap := map[string]any{
        "type":              typeName,
    }

    // Add the Data field to the map
    dataValue := reflect.ValueOf(p.data)
    if dataValue.Kind() == reflect.Ptr {
        dataValue = dataValue.Elem()
    }
    dataMap[string(typeName)] = dataValue.Interface()

    return json.Marshal(dataMap)
}

func NewTemplate(name, langCode string) *Template {
    return &Template{
        Name: name,
        Language: Language{
            Code: langCode,
        },
        Components: make([]Component, 0),
    }
}

func newComponent(tip ComponentType) *Component {
    return &Component{
        Type: tip,
        Parameters: make([]Parameter, 0),
    }
}

func NewHeaderComponent() *Component {
    return newComponent(ComponentTypeHeader)
}

func NewBodyComponent() *Component {
    return newComponent(ComponentTypeBody)
}

func (c *Component) AddParameter(m Message) *Component {
    c.Parameters = append(c.Parameters, Parameter{
        data: m,
    })
    return c
}

func (t *Template) AddComponent(c Component) *Template {
    t.Components = append(t.Components, c)
    return t
}

