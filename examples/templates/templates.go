package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/Lautarotetamusa/whatsapp-go/message"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")

    wa := whatsapp.NewWhatsapp(accessToken, numberId)
    recipient := "+54 341 15-585-4220" // put a valid number with country code

    template := message.NewTemplate("image_template", "es_AR").
        AddComponent(
            *message.NewHeaderComponent().
            AddParameter(&message.Image{ 
                ID: "903533311949481",
            }),
        )

    payload := whatsapp.NewPayload(template)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))

    res, err := wa.Send(recipient, template)

    if err != nil {
        fmt.Printf("%#v\n", err)
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
