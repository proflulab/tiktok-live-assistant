package configs

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"tiktok-live-assistant/models"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 将相对路径与当前工作目录结合，生成绝对路径
	dbPath := filepath.Join(cwd, "public", "db", "data.db")

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移，确保数据库中的表结构与模型匹配
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
