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
    
    msg := whatsapp.NewContacts()

    _, err := wa.Send(recipient, msg)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

func a(){
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    accessToken := os.Getenv("ACCESS_TOKEN")
    numberId    := os.Getenv("NUMBER_ID")
    recipient   := os.Getenv("RECIPIENT")  // valid number with country code

    wa := whatsapp.New(accessToken, numberId)

    msg := whatsapp.NewContacts(
        *whatsapp.NewContact(
            whatsapp.NewName("jose gonzales"), 
            // One contact can have multiple phone numbers
            whatsapp.NewPhone("+5493415854220"),
            whatsapp.NewPhone("+5493416668989"),
        ),
        *whatsapp.NewContact(
            // Specify first, middle, formmated name and other Name fields 
            whatsapp.Name{
                FormattedName: "pedro J. alberto",
                FirstName: "Pedro",
                LastName: "Alberto",
                MiddleName: "Jose",
            },
            // Phone with WaID have the Open Chat button
            whatsapp.Phone{
                Phone: "+5493418981233",
                WaID: "12345",
            },
        ),
    )

    payload := whatsapp.NewPayload(recipient, msg)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))

    res, err := wa.Send(recipient, msg)

    if err != nil {
        fmt.Println(err)
        panic("cannot send the whatsapp")
    }
    fmt.Printf("Message sended succesfully! Response %#v\n", res)
}
