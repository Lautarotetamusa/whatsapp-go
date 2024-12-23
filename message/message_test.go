package message

import "testing"

func TestImageCantHaveIdAndLink(t *testing.T) {
    msg := Image{
        ID: "12345",
        Link: "https://hola.com",
    }

    if err := msg.Validate(); err == nil {
        t.Fatalf("image with id and link must be an error")
    }
}

func TestVideoCantHaveIdAndLink(t *testing.T) {
    msg := Video{
        ID: "12345",
        Link: "https://hola.com",
    }

    if err := msg.Validate(); err == nil {
        t.Fatalf("video with id and link must be an error")
    }
}

func TestDocumentCantHaveIdAndLink(t *testing.T) {
    msg := Document{
        ID: "12345",
        Link: "http://test.com/",
    }

    if err := msg.Validate(); err == nil {
        t.Fatalf("document with id and link must be an error")
    }
}
