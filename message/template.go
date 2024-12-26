package message

import (
	"encoding/json"
	"fmt"
)

// https://developers.facebook.com/docs/whatsapp/cloud-api/guides/send-message-templates/
type Template struct {
	Name       string       `json:"name"`
    Language   Language     `json:"language"`
	Components []Component  `json:"components"`
}

type Language struct {
	Code LanguageCode `json:"code"`
}

func (t *Template) Validate() error {
    // TODO
    return nil
}

func (t *Template) GetType() MessageType {
    return TemplateType
}

type Parameter struct {
    data        Message
}

type Component struct {
    Type       ComponentType    `json:"type"`

	Parameters []Parameter      `json:"parameters"`
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
    if p.data == nil {
        return nil, fmt.Errorf("payload data its empty")
    }
    typ := p.data.GetType() 

    // Create a map to hold the payload data
    dataMap := map[string]any{
        "type": typ,
        // the Data field. eg. "video": p.data
        string(typ): p.data,
    }

    return json.Marshal(dataMap)
}

func NewTemplate(name string, langCode LanguageCode) *Template {
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

