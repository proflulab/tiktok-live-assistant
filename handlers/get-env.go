package handlers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetTiktokURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// 赋值
	return os.Getenv("TIKTOK_URL")
}

func GetTiktokLiveURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	liveURL := os.Getenv("TIKTOK_LIVE_URL")
	roomID := os.Getenv("ROOM_ID")
	return liveURL + roomID
}

func GetRoomXPATH() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("ROOM_XPATH")
}
