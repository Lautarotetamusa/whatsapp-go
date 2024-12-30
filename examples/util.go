package main

import (
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/joho/godotenv"
)

var (
    wa *whatsapp.Whatsapp
    to string
)

func getWhatsapp() {
    if wa != nil {
        return
    }

	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")
    to          = os.Getenv("RECIPIENT")  // valid number with country code

    wa = whatsapp.New(accessToken, numberId)
}
