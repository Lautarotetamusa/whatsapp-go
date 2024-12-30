package main

import (
	"fmt"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")
    to          := os.Getenv("RECIPIENT")  // valid number with country code

    wa := whatsapp.New(accessToken, numberId)
    msg := &whatsapp.Image{
        Media: whatsapp.NewMedia("903533311949481"),
    }

    payload := whatsapp.NewPayload(to, msg)
    fmt.Println(payload)
    res, err := wa.Send(to, msg)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Message sended successfully! Response %#v\n", res)
}
