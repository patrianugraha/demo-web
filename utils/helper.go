package utils

import (
	"os"
	"time"
)

// GetCurrentTime get current time ini DD-MMMM-YYYY HH:mm:ss format
func GetCurrentTime() string {
	currentTime := time.Now().Format("02-Jan-2006 15:04:05")

	return currentTime
}

// GetGreetingText get greeting text from config
func GetGreetingText() string {
	return os.Getenv("GREETING")
}

// GetBgColor get bg color from config
func GetBgColor() string {
	return os.Getenv("BG_COLOR")
}

// GetSecret get value from secret
func GetSecret() string {
	return os.Getenv("demo-web-secret")
}
