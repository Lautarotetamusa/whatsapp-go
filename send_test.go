package whatsapp

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const to = "+5493415854220"

func defaultWhatsapp() *Whatsapp {
	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
    }

    accessToken, ok := os.LookupEnv("ACCESS_TOKEN"); 
    if !ok || accessToken == "" {
        panic("ACCESS_TOKEN is required in the environment")
    }

    numberId, ok := os.LookupEnv("NUMBER_ID"); 
    if !ok || numberId == "" {
        panic("NUMBER_ID is required in the environment")
    }

    logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
    w := NewWhatsapp(accessToken, numberId, logger)
    return w
}

func TestSendText(t *testing.T){
    w := defaultWhatsapp()
    res, err := w.SendText(to, "hola")
    fmt.Printf("%v\n", res)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestSendTextWithUrl(t *testing.T){
    w := defaultWhatsapp()
    res, err := w.SendText(to, "https://www.youtube.com/watch?v=q1j7KygaRBo")
    fmt.Printf("%v\n", res)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}
