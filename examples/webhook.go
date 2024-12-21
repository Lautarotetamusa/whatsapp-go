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
        fmt.Printf("status changed %#v\n", s)
    })

    fmt.Println("Server running")

    http.ListenAndServe(":8080", webhook)
}
