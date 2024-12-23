package main

import (
	"fmt"
	"os"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/Lautarotetamusa/whatsapp-go/message"
)

func main(){
    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")

    wa := whatsapp.NewWhatsapp(accessToken, numberId)
    recipient := "+1234567899999" // put a valid number with country code

    res, err := wa.Send(recipient, &message.Text{
        Body: "hola!",
    })
    if err != nil {
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
