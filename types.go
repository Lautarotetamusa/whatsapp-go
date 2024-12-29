package whatsapp

// The payload data types must implement this interface
type Message interface {
    // Check if the message its valid
    Validate() error
    // Return the message type name. eg: "video"
    GetType() MessageType
}

// General Media type
type Media struct {
    // Only one of
    Link    string     `json:"link,omitempty"`
	ID      string     `json:"id,omitempty"`
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Image struct {
    *Media
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Video struct {
    *Media
    // Optional
    Caption string     `json:"caption,omitempty"`
}

type Document struct {
    *Media
    // Optional
	Filename string     `json:"filename,omitempty"`
	Caption  string     `json:"caption,omitempty"`
}

type Audio struct {
    *Media
}

type Sticker struct {
    *Media
}

type Contacts []Contact

// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/contacts-messages#post-body-parameters
type Contact struct {
	Name      Name      `json:"name"`
    // Contact's birthday in YYYY-MM-DD format.
	Birthday  string    `json:"birthday,omitempty"`
	Org       *Org       `json:"org,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
	Emails    []Email   `json:"emails,omitempty"`
	Phones    []Phone   `json:"phones,omitempty"`
	URLs      []URL     `json:"urls,omitempty"`
}

type Address struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
    // Postal or ZIP code
	Zip         string `json:"zip"`
	Country     string `json:"country"`
    // ISO two-letter country code.
	CountryCode string `json:"country_code"`
    // Type of address, such as home or work.
	Type        string `json:"type"`
}

type Email struct {
    // Email address of the contact.
	Email string `json:"email"`
    // Type of email, such as personal or work.
	Type  string `json:"type"`
}

type Name struct {
    // Contact's formatted name. This will appear in the message alongside the profile arrow button.
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

type Org struct {
	Company    string `json:"company"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type Phone struct {
    // This is optionall in the documentation (??
	Phone string `json:"phone"`
    // Type of phone number. For example, cell, mobile, main, iPhone, home, work, etc.
	Type  string `json:"type,omitempty"`
    // If omitted, the message will display an Invite to WhatsApp button instead of the standard buttons.
    // see: https://developers.facebook.com/docs/whatsapp/cloud-api/messages/contacts-messages#button-behavior
	WaID  string `json:"wa_id,omitempty"`
}

type URL struct {
    // Website URL associated with the contact or their company.
	URL  string `json:"url"`
    // Type of website. For example, company, work, personal, Facebook Page, Instagram, etc.
	Type string `json:"type"`
}
