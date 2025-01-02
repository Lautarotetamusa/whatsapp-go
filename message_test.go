package whatsapp

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMediaCantHaveIdAndLink(t *testing.T) {
	msg := &Image{
        Media: &Media{
            ID:   "12345",
            Link: "https://hola.com",
        },
	}

    expectValidationErr(t, msg, ErrorIdAndLink.Error())
}

func TestText(t *testing.T) {
	msg := &Text{
		Body: "",
	}
    expectValidationErr(t, msg, "body cannot be empty")

	msg = NewTextMessage("http://httpbin.org/")

	if !msg.PreviewUrl {
		t.Errorf("new text message must have a preview url")
	}
}

func TestNewName(t *testing.T) {
	msg := NewName("juan pablo")

	if err := msg.Validate(); err != nil {
		t.Errorf("NewName doesnt construct a valid Name")
	}
}

func TestContactName(t *testing.T) {
	name := Name{
		FormattedName: "juan",
	}
	msg := NewContacts(*NewContact(name, Phone{}))

    expectValidationErr(t, msg, "first_name, last_name and formmatted_name are required")
}

func TestContacts(t *testing.T) {
	msg := NewContacts()
    expectValidationErr(t, msg, "you need to send at least one contact")

	msg = NewContacts(*NewContact(NewName("juan pablo")))
    expectValidationErr(t, msg, "contact must have at least one phone")
}

func TestContactMustHaveName(t *testing.T) {
	contact := Contact{}
	contact.Phones = []Phone{NewPhone("12345")}
    msg := NewContacts(contact)
    
    expectValidationErr(t, msg, "first_name, last_name and formmatted_name are required")
}

func TestButtons(t *testing.T) {
	msg := NewInteractive(NewButtons(
		NewButton("btn_1", "111"),
		NewButton("btn_2", "112"),
		NewButton("btn_3", "113"),
		NewButton("btn_4", "114"),
	)).SetBody("body")

    expectValidationErr(t, msg, "interaction cannot have more than 3 buttons")

	msg = NewInteractive(NewButtons(
		NewButton("btn_1", "111"),
		NewButton("btn_2", "111"),
	)).SetBody("body")

    expectValidationErr(t, msg, "two buttons cannot have the same id in one message")
}

func TestButtonID(t *testing.T) {
	btn := NewButton("title", "")
	btn.Reply.ID = strings.Repeat("A", 256)
    msg := NewInteractive(NewButtons(btn)).SetBody("body")

    expectNoValidationErr(t, msg)

	btn.Reply.ID += "A"
    msg = NewInteractive(NewButtons(btn)).SetBody("body")
    expectValidationErr(t, msg, "button id cannot have more than 256 characters")

	btn.Reply.ID = " 1234 "
    msg = NewInteractive(NewButtons(btn)).SetBody("body")
    expectValidationErr(t, msg, "button id cannot start or end with a space")
}

func TestButtonTitle(t *testing.T) {
	btn := NewButton("", "114")
    msg := NewInteractive(NewButtons(btn)).SetBody("body")
    expectValidationErr(t, msg, "button title cannot be empty")

	btn.Reply.Title = strings.Repeat("1", 20)
    msg = NewInteractive(NewButtons(btn)).SetBody("body")
    expectNoValidationErr(t, msg)

    btn.Reply.Title += "1"
    msg = NewInteractive(NewButtons(btn)).SetBody("body")
    expectValidationErr(t, msg, "button title cannot have more than 20 characters")
}

func TestButtonMessageMustHaveBody(t *testing.T) {
    btn := NewInteractive(NewButtons(NewButton("Click me!", "1")))
    expectValidationErr(t, btn, "body its required")
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

func buildListFromRow(row Row) *Interactive {
    return NewInteractive(
        NewList("button", 
            NewListSection("Section 1", row),
        ),
    ).
    SetBody("body")
}


func expectNoValidationErr(t *testing.T, msg Message) {
    payload := NewPayload(to, msg)

    if err := payload.Validate(); err != nil { // error
        t.Errorf("expected no error, received: %s", err.Error())
	}
}
