package handlers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(envName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if envName == "TIKTOK_LIVE_URL" {
		liveURL := os.Getenv(envName)
		roomID := os.Getenv("ROOM_ID")
		return liveURL + roomID
	} else {
		return os.Getenv(envName)
	}
}
