package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go/webhook"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    verifyToken := os.Getenv("VERIFY_TOKEN")

    wh := webhook.New(verifyToken)

    wh.OnStatusChange(func (s *webhook.Status){
        fmt.Printf("status changed: %#v\n", s)
    })

    wh.OnNewMessage(func (m *webhook.Message) {
        fmt.Printf("new message received: %#v\n", m)
    })

    fmt.Println("Server running")

    http.ListenAndServe(":3000", wh)
}
