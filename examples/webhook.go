package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    verifyToken := os.Getenv("VERIFY_TOKEN")

    webhook := whatsapp.NewWebhook(verifyToken)
    webhook.OnStatusChange(func (s *whatsapp.Status){
        fmt.Printf("status changed: %#v\n", s)
    })
    webhook.OnNewMessage(func (m *whatsapp.Message) {
        fmt.Printf("new message recived: %#v\n", m)
    })

    fmt.Println("Server running")

    http.ListenAndServe(":3000", webhook)
}
