package main

import (
	"tiktok-live-assistant/configs"
	"tiktok-live-assistant/services"
)

// 主程序入口
func main() {
	// 初始化数据库配置
	configs.InitDB()
	// 开启浏览器操作服务
	services.ChromeCtrl()
}
