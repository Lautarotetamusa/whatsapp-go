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

    template := message.NewTemplate("image_template", message.SpanishARG).
        AddComponent(
            *message.NewHeaderComponent().
            AddParameter(&message.Image{ 
                Media: message.FromID("903533311949481"),
            }),
        )

    // Maybe this style its better?
    // t := message.NewTemplate("image_template", "es_AR", 
    //     NewHeaderComponent(
    //          &message.Image{ 
    //              Media: message.FromID("903533311949481"),
    //          }
    //     ),
    //     NewBodyComponent(
    //         &message.Text{}
    //     )
    // ).

    payload := whatsapp.NewPayload(recipient, template)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))

    res, err := wa.Send(recipient, template)

    if err != nil {
        fmt.Println(err)
        panic("cannot send the message")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
