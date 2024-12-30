package whatsapp

import (
	"strings"
	"testing"
)

func TestMediaCantHaveIdAndLink(t *testing.T) {
	msg := &Media{
		ID:   "12345",
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

	if !msg.PreviewUrl {
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

func TestButtons(t *testing.T) {
	btn := NewButtons(
		NewButton("btn_1", "111"),
		NewButton("btn_2", "112"),
		NewButton("btn_3", "113"),
		NewButton("btn_4", "114"),
	)

	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("buttons interactive action can have max of 3 buttons")
		}
	}

	btn = NewButtons(
		NewButton("btn_1", "111"),
		NewButton("btn_2", "111"),
	)

	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("two buttons with the same id must be an error")
		}
	}
}

func TestButtonID(t *testing.T) {
	id := strings.Repeat("A", 256)
	btn := NewButton("12345678901234567890", id)
	if err := btn.Validate(); err != nil {
		t.Error("button id can have 256 characters")
	}

	id = strings.Repeat("A", 257)
	btn = NewButton("12345678901234567890", id)
	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("button id cannot have more than 256 characters")
		}
	}

	id = " 1234 "
	btn = NewButton("12345678901234567890", id)
	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("button id cannot start or end with a space")
		}
	}
}

func TestButtonTitle(t *testing.T) {
	btn := NewButton("", "114")

	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("button title cannot be empty")
		}
	}

	btn = NewButton("12345678901234567890", "114")
	if err := btn.Validate(); err != nil {
		t.Error("button title can have exactly 20 characters")
	}

	btn = NewButton("123456789012345678901", "114")
	if err := btn.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("button title cannot have more than 20 characters")
		}
	}
}
