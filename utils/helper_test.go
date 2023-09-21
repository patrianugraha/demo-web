package utils

import (
	"strings"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func TestGetGreetingText(t *testing.T) {
	godotenv.Load("../.env")
	greetingText := GetGreetingText()
	if strings.TrimSpace(greetingText) == "" {
		t.Error("Greeting text is empty")
	}
}

func TestGetCurrentTime(t *testing.T) {
	godotenv.Load("../.env")
	currentTime := GetCurrentTime()
	if strings.TrimSpace(currentTime) == "" {
		t.Error("Current time text is empty")
	}
}

func TestGetBgColor(t *testing.T) {
	godotenv.Load("../.env")
	bgColor := GetBgColor()
	if strings.TrimSpace(bgColor) == "" {
		t.Error("Bg Color text is empty")
	}
}
