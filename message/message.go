package message

import (
	"errors"
	"strings"
)

func isUrl(s string) bool {
	return strings.Contains(s, "http://") || strings.Contains(s, "https://")
}

func NewTextMessage(msg string) *Text {
	return &Text{
        // if the message its an url make the message shows the preview
		PreviewUrl: isUrl(msg),
		Body:       msg,
	}
}

// Construct new contacts message
func NewContacts(contacts ...Contact) *Contacts {
    c := Contacts(contacts)
	return &c
}

// Add appends new contacts to the collection
func (c *Contacts) Add(contacts ...Contact) *Contacts {
	*c = append(*c, contacts...)
    return c
}

// name := NewName("juan carlos")
// Get the full name and returns a Name object
// The FullName must have at least ONE space "firstName lastName"
func NewName(fullName string) Name {
    var name Name
    splitted := strings.Split(fullName, " ")

    if len(splitted) > 1 {
        name.FormattedName = fullName
        name.FirstName = splitted[0]
        name.LastName = splitted[1]
    }
    return name
}

func NewPhone(phone string) Phone {
    return Phone{
        Phone: phone,
    }
}

func NewPhoneWithWaID(phone, waid string) Phone {
    return Phone{
        Phone: phone,
        WaID: waid,
    }
}

func NewContact(name Name, phones ...Phone) *Contact {
    return &Contact{
        Name: name,
        Phones: phones,
    }
}

func (c *Contact) AddAdress(a Address) *Contact {
    c.Addresses = append(c.Addresses, a)
    return c
}

func (c *Contact) AddPhone(phone Phone) *Contact {
    c.Phones = append(c.Phones, phone)
    return c
}

func (c *Contact) AddUrl(url URL) *Contact {
    c.URLs = append(c.URLs, url)
    return c
}

func (c *Contact) AddEmail(email Email) *Contact {
    c.Emails = append(c.Emails, email)
    return c
}

func (c *Contact) SetOrg(org Org) *Contact {
    c.Org = &org
    return c
}

func (c *Contact) SetBirthday(birthday string) *Contact {
    c.Birthday = birthday
    return c
}

func (c *Contact) SetName(name Name) *Contact {
    c.Name = name
    return c
}

// If the parameter its a url set the Link otherwise set the ID
func NewMedia(idOrLink string) *Media {
    m := Media{}
    if isUrl(idOrLink){
        m.Link = idOrLink
    }else{
        m.ID = idOrLink
    }
    return &m
}

func FromID(id string) *Media {
    return &Media{
        ID: id,
    }
}

func FromLink(link string) *Media {
    return &Media{
        Link: link,
    }
}

func (n *Name) Validate() error {
    if n.FirstName == "" || n.LastName == "" || n.FormattedName == "" {
        return NewErr(&Contacts{}, errors.New("first_name, last_name and formmatted_name are required"))
    }
    return nil
}

func (c *Contact) Validate() error {
    if len(c.Phones) == 0 {
        return NewErr(&Contacts{}, errors.New("contact must have at least one phone"))
    }
    return c.Name.Validate()
}

func (c *Contacts) Validate() error {
    if len(*c) == 0 {
        return NewErr(c, errors.New("you need to send at least one contact"))
    }
    for _, contact := range *c {
        if err := contact.Validate(); err != nil {
            return err
        }
    }
    return nil
}

func (m *Media) Validate() error {
    if (m.Link != "") && (m.ID != "") {
        return ErrorIdAndLink
    }
    if (m.Link == "") && (m.ID == "") {
        return ErrorIdAndLink
    }
    return nil
}

func (m *Text) Validate() error { 
    if m.Body == "" {
        return NewErr(m, errors.New("body cannot be empty"))
    }
    return nil
}

// implements the Message interface
// Media types
func (m *Image)     GetType() MessageType { return ImageType }
func (m *Video)     GetType() MessageType { return VideoType }
func (m *Audio)     GetType() MessageType { return AudioType }
func (m *Sticker)   GetType() MessageType { return StickerType }
func (m *Document)  GetType() MessageType { return DocumentType }
func (m *Text)      GetType() MessageType { return TextType }

func (m *Contacts)  GetType() MessageType { return ContactsType }
