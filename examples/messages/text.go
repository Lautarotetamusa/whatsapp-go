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

    wa := whatsapp.New(accessToken, numberId)
    recipient := "+54 341 15-585-4220" // put a valid number with country code

    res, err := wa.Send(recipient, &whatsapp.Text{
        Body: "hola!",
    })

    if err != nil {
        fmt.Printf("%#v\n", err)
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
