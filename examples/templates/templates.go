package main

import (
	"encoding/json"
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
    recipient   := os.Getenv("RECIPIENT")  // valid number with country code

    wa := whatsapp.New(accessToken, numberId)

    template := whatsapp.NewTemplate("image_template", whatsapp.SpanishARG).
        AddComponent(
            *whatsapp.NewHeaderComponent().
            AddParameter(&whatsapp.Image{ 
                Media: whatsapp.FromID("903533311949481"),
            }),
        )

    // Maybe this style its better?
    // t := whatsapp.NewTemplate("image_template", "es_AR", 
    //     NewHeaderComponent(
    //          &whatsapp.Image{ 
    //              Media: whatsapp.FromID("903533311949481"),
    //          }
    //     ),
    //     NewBodyComponent(
    //         &whatsapp.Text{}
    //     )
    // ).

    payload := whatsapp.NewPayload(recipient, template)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))

    res, err := wa.Send(recipient, template)

    if err != nil {
        fmt.Println(err)
        panic("cannot send the whatsapp")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
