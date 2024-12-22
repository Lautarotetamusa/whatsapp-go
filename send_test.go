package whatsapp

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const to = "54341155854220"

var wa *Whatsapp

func getWhatsapp() *Whatsapp {
    if wa != nil {
        return wa
    }

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

    wa = NewWhatsapp(accessToken, numberId)    
    return wa
}

func TestSendText(t *testing.T){
    w := getWhatsapp()
    _, err := w.SendText(to, "hola")
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestSendTextWithUrl(t *testing.T){
    w := getWhatsapp()
    _, err := w.SendText(to, "https://www.youtube.com/watch?v=q1j7KygaRBo")
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestSendImageWithId(t *testing.T) {
    w := getWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    imageId := "903533311949481"

    payload := MediaPayload{
        ID: imageId,
        Caption: "Test image",
    }

    _, err := w.SendImage(to, payload)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestSendVideoWithId(t *testing.T) {
    w := getWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    imageId := "967038865474431"

    payload := MediaPayload{
        ID: imageId,
        Caption: "Test video",
    }

    _, err := w.SendVideo(to, payload)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestMediaCantHaveIdAndLink(t *testing.T) {
    w := getWhatsapp()

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
    w := getWhatsapp()
    //TODO: this id lasts 15 days, the test will fail
    docId := "903533311949481"

    payload := DocumentPayload{
        ID: docId,
        Caption: "Test document",
        Filename: "test_file.pdf",
    }

    _, err := w.SendDocument(to, payload)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
}

func TestDocumentCantHaveIdAndLink(t *testing.T) {
    w := getWhatsapp()

    payload := DocumentPayload{
        ID: "12345",
        Link: "http://test.com/",
    }

    _, err := w.SendDocument(to, payload)
    if err == nil {
        t.Fatalf("media with id and link must be an error")
    }
}
