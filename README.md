[![Go](https://github.com/Lautarotetamusa/whatsapp-go/actions/workflows/go.yml/badge.svg)](https://github.com/Lautarotetamusa/whatsapp-go/actions/workflows/go.yml)
# whatsapp cloud api

Simple and easy to use wraper for the meta cloud api \
Easy way for creating bots and sending and reciving messages using the *Whatsapp Cloud API* \

# Supported API
- [ ] Message
    - [x] Text
    - [x] Media
        - [x] Video
        - [x] Image
        - [x] Audio
        - [x] Sticker
        - [x] Document
    - [ ] Interactive
        - [x] Buttons
        - [ ] Call To Action
        - [ ] Catalog
        - [ ] List
        - [ ] Product
        - [ ] Product List
        - [ ] Flow
    - [ ] Template
- [x] Webhooks
    - [x] Reciving messages
- [ ] Media management

# Instalation
`go get github.com/Lautarotetamusa/whatsapp-go`

# Message
## Simple text message
```go
package main

import (
	"fmt"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/Lautarotetamusa/whatsapp-go/message"
)

func main(){
    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")

    wa := whatsapp.New(accessToken, numberId)
    recipient := "+1234567899999" // put a valid number with country code

    res, err := wa.Send(recipient, &message.Text{
        Body: "hola!",
    })
    if err != nil {
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
```

## Send Media
```go
res, err := wa.Send("recipient-phone", &whatsapp.Image{
    Media: whatsapp.FromID("valid-image-id"),
    Caption: "Test image",
})
```
or with a link
```go
res, err := wa.Send("recipient-phone", &whatsapp.Video{
    Media: whatsapp.FromLink("https://url.com/image"),
})
```

## Send Contacts
```go
msg := whatsapp.NewContacts(
    *whatsapp.NewContact(
        // NewName expects string with at least one space
        // "FirstName LastName"
        whatsapp.NewName("jose gonzales"), 
        // One contact can have multiple phone numbers
        whatsapp.NewPhone("+5493415854220"),
        whatsapp.NewPhone("+5493416668989"),
    ),
    *whatsapp.NewContact(
        // Specify first, middle, formmated name and other Name fields 
        whatsapp.Name{
            FormattedName: "pedro J. alberto",
            FirstName: "Pedro",
            LastName: "Alberto",
            MiddleName: "Jose",
        },
        // Phone with WaID will have the Open Chat button
        whatsapp.Phone{
            Phone: "+5493418981233",
            WaID: "12345",
        },
    ),
)
res, err := wa.Send(to, msg)
```

## Send Interactive 
### Buttons
```go
msg := NewInteractive(NewButtons(
    NewButton("btn_1", "123"), // btn_name, btn_id
    NewButton("btn_2", "124"),
)).
SetBody("hi!").
SetHeader(&Image{
    Media: FromID("valid-image-id"),
})

res, err := wa.Send(to, msg)
```

# Webhook
```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go/webhook"
)

func main(){
    verifyToken := os.Getenv("VERIFY_TOKEN")

    webhook := webhook.New(verifyToken)
    webhook.OnStatusChange(func (s *whatsapp.Status){
        fmt.Printf("status changed: %#v\n", s)
    })
    webhook.OnNewMessage(func (m *whatsapp.Message) {
        fmt.Printf("new message recived: %#v\n", m)
    })

    fmt.Println("Server running")

    http.ListenAndServe(":3000", webhook)
}
```

# Documentation
[Webhooks](https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components/) \
[Messages](https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages)
