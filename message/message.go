package message

import (
	"fmt"
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
func (c *Contacts) Add(contacts ...Contact) {
	*c = append(*c, contacts...)
}

// c := NewContact("juan", Phone{
//     Phone: "+5493148921"
//     WaID: "19175559999"
// )}
func NewContact(name string, phones ...Phone) Contact {
    return Contact{
        Name: Name{
            FormattedName: name,
        },
        Phones: phones,
    }
}

func (c *Contact) AddAdress(a Address) {
    c.Addresses = append(c.Addresses, a)
}

func (c *Contact) AddPhone(phone Phone) {
    c.Phones = append(c.Phones, phone)
}

func (c *Contact) AddUrl(url URL) {
    c.URLs = append(c.URLs, url)
}

func (c *Contact) AddEmail(email Email) {
    c.Emails = append(c.Emails, email)
}

func (c *Contact) SetOrg(org Org) {
    c.Org = &org
}

func (c *Contact) SetBirthday(birthday string) {
    c.Birthday = birthday
}

func (c *Contact) SetName(name Name) {
    c.Name = name
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

func (c *Contacts) Validate() error {
    // TODO
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
        return fmt.Errorf("text body cannot be empty")
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
