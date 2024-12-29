package whatsapp

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// For test numbers, the recipient must be on a list of available phones
const to = "+54 341 15-585-4220"

var wa *Whatsapp

func getWhatsapp() *Whatsapp {
	if wa != nil {
		return wa
	}

	if err := godotenv.Load("./.env"); err != nil {
		panic("Unable to load .env file")
	}

	accessToken, ok := os.LookupEnv("ACCESS_TOKEN")
	if !ok || accessToken == "" {
		panic("ACCESS_TOKEN is required in the environment")
	}

	numberId, ok := os.LookupEnv("NUMBER_ID")
	if !ok || numberId == "" {
		panic("NUMBER_ID is required in the environment")
	}

	wa = New(accessToken, numberId)
	return wa
}

func TestSendText(t *testing.T) {
	w := getWhatsapp()
	_, err := w.SendText(to, "hola")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendTextWithUrl(t *testing.T) {
	w := getWhatsapp()
	_, err := w.SendText(to, "https://www.youtube.com/watch?v=q1j7KygaRBo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendImageWithId(t *testing.T) {
	w := getWhatsapp()
	//TODO: this id lasts 15 days, the test will fail
	imageId := "903533311949481"

	msg := Image{
		Media:   FromID(imageId),
		Caption: "Test image",
	}

	_, err := w.Send(to, &msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoWithId(t *testing.T) {
	w := getWhatsapp()
	//TODO: this id lasts 15 days, the test will fail
	videoId := "967038865474431"

	msg := Video{
		Media:   FromID(videoId),
		Caption: "Test video",
	}

	_, err := w.Send(to, &msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocumentWithId(t *testing.T) {
	w := getWhatsapp()
	//TODO: this id lasts 15 days, the test will fail
	docId := "903533311949481"

	msg := Document{
		Media:    FromID(docId),
		Caption:  "Test document",
		Filename: "test_file.pdf",
	}

	_, err := w.Send(to, &msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioWithId(t *testing.T) {
	w := getWhatsapp()
	//TODO: this id lasts 15 days, the test will fail
	audioId := "903533311949481"

	msg := Audio{
		Media: FromID(audioId),
	}

	_, err := w.Send(to, &msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendStickerWithId(t *testing.T) {
	w := getWhatsapp()
	//TODO: this id lasts 15 days, the test will fail
	stickerId := "903533311949481"

	msg := Sticker{
		Media: FromID(stickerId),
	}

	_, err := w.Send(to, &msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendContacts(t *testing.T) {
	w := getWhatsapp()
	msg := NewContacts(
		*NewContact(
			NewName("jose gonzales"),
			// One contact can have multiple phone numbers
			NewPhone("+5493415854220"),
			NewPhone("+5493416668989"),
		),
	)

	_, err := w.Send(to, msg)
	if err != nil {
		t.Fatal(err)
	}
}
