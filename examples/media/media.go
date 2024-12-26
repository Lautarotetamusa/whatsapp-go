package main

import (
	"encoding/json"
	"fmt"

	"github.com/Lautarotetamusa/whatsapp-go"
	"github.com/Lautarotetamusa/whatsapp-go/message"
)

func main(){
    m := &message.Video{
        ID: "valid-id",
    }

    payload := whatsapp.NewPayload(m)
    t, err := json.MarshalIndent(payload, "", "   ")
    fmt.Println(err)
    fmt.Printf("%s\n", string(t))
}
