[![Go](https://github.com/Lautarotetamusa/whatsapp-go/actions/workflows/go.yml/badge.svg)](https://github.com/Lautarotetamusa/whatsapp-go/actions/workflows/go.yml)
# whatsapp cloud api

Simple wraper for the meta cloud api

# Usage
## Instalation
`go get github.com/Lautarotetamusa/whatsapp-go`

## Send message
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

### Sending a image
```go
res, err := wa.Send("recipient-phone", &message.Image{
    ID: "valid-image-id",
    Caption: "Test image",
})
```

## Webhook
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
[Webhooks](https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components/)
[Messages](https://developers.facebook.com/docs/whatsapp/cloud-api/messages/)
