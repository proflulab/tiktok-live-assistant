package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// 主程序入口
func main() {

	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 赋值
	tiktokURL := os.Getenv("TIKTOK_URL")
	liveURL := os.Getenv("TIKTOK_LIVE_URL")
	roomId := os.Getenv("ROOM_ID")

	fmt.Println("TIKTOK_URL:\n", tiktokURL)
	fmt.Println("TIKTOK_LIVE_URL:\n", liveURL)
	fmt.Println("ROOM_ID:\n", roomId)

}
