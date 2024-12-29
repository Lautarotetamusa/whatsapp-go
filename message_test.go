package whatsapp

import "testing"

func TestMediaCantHaveIdAndLink(t *testing.T) {
    msg := &Media{
        ID: "12345",
        Link: "https://hola.com",
    }

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("media with id and link must be a ValidationError")
        }
    }
}

func TestTextCantHaveEmptyBody(t *testing.T) {
    msg := &Text{
        Body: "",
    }

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("text with empty body must be a ValidationError")
        }
    }
}

func TestNewMessageWithUrl(t *testing.T) {
    msg := NewTextMessage("http://httpbin.org/")

    if !msg.PreviewUrl  {
        t.Fatalf("new text message must have a preview url")
    }
}

func TestNewName(t *testing.T) {
    msg := NewName("juan pablo")

    if err := msg.Validate(); err != nil {
        t.Fatalf("NewName doesnt construct a valid Name")
    }
}

func TestNameWithNoFirstName(t *testing.T) {
    msg := Name{
        FormattedName: "juan",
    }

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("name with only formatted_name must be a ValidationError")
        }
    }
}

func TestAtLeastOneContact(t *testing.T) {
    msg := NewContacts()

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("contacts with no contact must be a ValidationError")
        }
    }
}

func TestContactMustHaveAtLeastOnePhone(t *testing.T) {
    msg := NewContact(NewName("juan pablo"))

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("contact with no phone must be a ValidationError")
        }
    }
}

func TestContactMustHaveName(t *testing.T) {
    msg := Contact{}
    msg.Phones = []Phone{NewPhone("12345")}

    if err := msg.Validate(); err == nil {
        if _, ok := err.(ValidationError); !ok {
            t.Fatalf("contact with no name must be a ValidationError")
        }
    }
}
