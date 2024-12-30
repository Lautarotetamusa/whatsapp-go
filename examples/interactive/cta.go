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

    msg := whatsapp.NewCallToAction("hola!", "See Dates", "https://www.luckyshrub.com").
        SetHeader(&whatsapp.Image{
            Media: whatsapp.NewMedia("903533311949481"),
        }).
        SetFooter("lautarotetamusa")

    wa := whatsapp.New(accessToken, numberId)

    // payload := whatsapp.NewPayload(to, msg)
    // fmt.Println(payload)

    res, err := wa.Send(to, msg)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}