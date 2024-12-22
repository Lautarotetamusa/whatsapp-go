package whatsapp

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const to = "54341155854220"

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
    // fmt.Println(accessToken)
    // fmt.Println(numberId)

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

func TestSendImageWithId(t *testing.T) {
    w := defaultWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    imageId := "903533311949481"

    payload := MediaPayload{
        ID: imageId,
        Caption: "Test image",
    }

    res, err := w.SendImage(to, payload)

    fmt.Printf("%v\n", res)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestSendVideoWithId(t *testing.T) {
    w := defaultWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    imageId := "967038865474431"

    payload := MediaPayload{
        ID: imageId,
        Caption: "Test video",
    }

    res, err := w.SendVideo(to, payload)

    fmt.Printf("%v\n", res)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestMediaCantHaveIdAndLink(t *testing.T) {
    w := defaultWhatsapp()

    imageId := "12345"
    payload := MediaPayload{
        ID: imageId,
        Link: "https://hola.com",
    }

    _, err := w.SendImage(to, payload)
    if err == nil {
        t.Fatalf("media with id and link must be an error")
    }
}

func TestSendDocument(t *testing.T) {
    w := defaultWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    docId := "903533311949481"

    payload := DocumentPayload{
        ID: docId,
        Caption: "Test document",
        Filename: "test_file.pdf",
    }

    res, err := w.SendDocument(to, payload)

    fmt.Printf("%v\n", res)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestDocumentCantHaveIdAndLink(t *testing.T) {
    w := defaultWhatsapp()

    payload := DocumentPayload{
        ID: "12345",
        Link: "http://test.com/",
    }

    _, err := w.SendDocument(to, payload)
    if err == nil {
        t.Fatalf("media with id and link must be an error")
    }
}
