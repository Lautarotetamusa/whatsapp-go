package main

import (
	"encoding/json"
	"fmt"

	"github.com/Lautarotetamusa/whatsapp-go"
)

func main(){
    m := &whatsapp.Video{
        Media: whatsapp.FromID("123"),
    }
    // valid phone number
    to := "12345678900"

    payload := whatsapp.NewPayload(to, m)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))
}
