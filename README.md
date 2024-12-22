# whatsapp cloud api

Simple wraper for the meta cloud api

# Usage
`go get github.com/Lautarotetamusa/whatsapp-go`

# Send message
```go
func main(){
    accessToken := os.GetEnv("ACCESS_TOKEN")
    numberId    := os.GetEnv("NUMBER_ID")

    wa := whatsapp.NewWhatsapp(accessToken, numberId)
    recipient = "+1234567899999" // put a valid number with country code

    res, err := wa.SendText(recipient, "hello!")
    if err != nil {
        panic("cannot send the message")
    }
}
```

# Webhook
```go
func main(){
    verifyToken := os.Getenv("VERIFY_TOKEN")

    webhook := whatsapp.NewWebhook(verifyToken)
    webhook.OnStatusChange(func (s *whatsapp.Status){
        fmt.Printf("status changed %#v\n", s)
    })

    http.ListenAndServe(":8080", webhook)
}
```

# Documentation
[Webhooks](https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components/)
