package handlers

import (
	"fmt"
	"os"
	"testing"
)

// func GenerateUniqueURL(settings []Setting) string {

func TestGenerateUniqueURL(t *testing.T) {
	settings := []Setting{{Label: "Form Name", Type: "text", Default: "Contact", Required: true},
		{Label: "Website", Type: "text", Default: "Umegain", Required: true},
		{Label: "ChannelID", Type: "text", Default: "0192dd70-cdf1-7e15-8776-4fee4a78405e", Required: true},
	}
	result := GenerateUniqueURL(settings)

	BASE_URL := os.Getenv("BASE_URL")

	expected := fmt.Sprintf("%s/formify/Umegain/0192dd70-cdf1-7e15-8776-4fee4a78405e", BASE_URL)

	if result != expected {
		t.Errorf("GenerateUniqueURL(settings) = %s; want %s", result, expected)
	}
}

func TestExtractText(t *testing.T) {
	msg := "<p>some msg</p><p></p>"
	result := ExtractText(msg)

	expected := "some msg"

	if result != expected {
		t.Errorf("ExtractText('<p>some msg</p><p></p>') = %s; want %s", result, expected)
	}
}
