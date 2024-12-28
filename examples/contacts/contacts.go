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
    recipient   := os.Getenv("RECIPIENT")  // valid number with country code

    wa := whatsapp.NewWhatsapp(accessToken, numberId)

    msg := message.NewContacts(
        message.NewContact("juan", message.Phone{
            Phone: "+5493415854220",
        }),
        message.NewContact("pedro", message.Phone{
            Phone: "+5493415854220",
        }),
    )

    payload := whatsapp.NewPayload(recipient, msg)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))

    res, err := wa.Send(recipient, msg)

    if err != nil {
        fmt.Println(err)
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
