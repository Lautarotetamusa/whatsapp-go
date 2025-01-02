package whatsapp

import (
	"fmt"
	"strconv"
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

func TestButtonMessageMustHaveBody(t *testing.T) {
    btn := NewInteractive(NewButtons(NewButton("Click me!", "1")))

    payload := NewPayload(to, btn)

    if err := payload.Validate(); err == nil {
		if _, ok := err.(ValidationError); !ok {
			t.Error("Interactive button with no body must be an error")
		}
	}
}

func expectValidationErr(t *testing.T, msg Message, errStr string) {
    payload := NewPayload(to, msg)

    if err := payload.Validate(); err == nil { // no error
        t.Errorf("expected error: %s. received no error", errStr)
	}else{
        if e, ok := err.(*ValidationError); ok {
            if e.err.Error() != errStr {
                t.Errorf("expected: %s. received: %s", errStr, e.err.Error())
            }
        }else{
            t.Error("expect ValidationError but found other error", err)
        }
    }
}

func expectNoValidationErr(t *testing.T, msg Message) {
    payload := NewPayload(to, msg)

    if err := payload.Validate(); err != nil { // error
        t.Errorf("expected no error, received: %s", err.Error())
	}
}

func TestListMessage(t *testing.T) {
    list := NewList(strings.Repeat("A", 21),
        NewListSection("Section 1", NewRow("1", "row 1", "desc 1")),
    )

    msg := NewInteractive(list).SetBody("body") 
    expectValidationErr(t, msg, "list button cannot have more than 20 chars")

    list.Button = strings.Repeat("A", 20)
    msg.action = list    
    expectNoValidationErr(t, msg)

    msg.action = NewList("button")
    expectValidationErr(t, msg, "list must have at least 1 section")

    list = NewList("button")
    for i := range 10 {
        list.Sections = append(list.Sections, NewListSection(fmt.Sprintf("section %d", i)))
    }
    msg.action = list
    expectNoValidationErr(t, msg)

    list.Sections = append(list.Sections, NewListSection("section 11"))
    msg.action = list
    expectValidationErr(t, msg, "list cannot have more than 10 sections")

    rows := []Row{}
    for i := range 10 {
        rows = append(rows, NewRow(strconv.Itoa(i), fmt.Sprintf("row - %d", i), fmt.Sprintf("desc %d", i)))
    }
    msg.action = NewList("button", ListSection{Rows: rows, Title: "button"})
    expectNoValidationErr(t, msg)

    rows = append(rows, NewRow("11", "row - 11", "desc 11"))
    msg.action = NewList("button", ListSection{Rows: rows, Title: "button"})
    expectValidationErr(t, msg, "list section cannot have more than 10 rows")
}

func buildListFromRow(row Row) *Interactive {
    return NewInteractive(
        NewList("button", 
            NewListSection("Section 1", row),
        ),
    ).
    SetBody("body")
}

func TestListMessageRows(t *testing.T) {
    // ID
    row := NewRow(strings.Repeat("1", 200), "title", "desc")
    expectNoValidationErr(t, buildListFromRow(row))

    row.ID += "1"
    expectValidationErr(t, buildListFromRow(row), "row id cannot have more than 200 characters")

    // Title
    row = NewRow("1", strings.Repeat("1", 24), "desc")
    expectNoValidationErr(t, buildListFromRow(row))

    row.Title += "1"
    expectValidationErr(t, buildListFromRow(row), "row title cannot have more than 24 characters")

    // Description
    row = NewRow("1", "title", strings.Repeat("A", 72))
    expectNoValidationErr(t, buildListFromRow(row))

    row.Description += "A"
    expectValidationErr(t, buildListFromRow(row), "row description cannot have more than 72 characters")

    row.Description = ""
    expectNoValidationErr(t, buildListFromRow(row))
}
