package main

import (
	"tiktok-live-assistant/configs"
	"tiktok-live-assistant/services"
)

// 主程序入口
func main() {

	configs.InitDB()

	services.ChromeCtrl()
}

//func test() {
//	for {
//		fmt.Println("请输入句子，或输入退出结束测试")
//		var sentence string
//		_, err := fmt.Scanln(&sentence)
//		if err != nil {
//			return
//		}
//		if sentence == "退出" {
//			break
//		}
//		flag := handlers.SentenceClassify(sentence)
//		if flag {
//			fmt.Println("结果为TRUE")
//		} else {
//			fmt.Println("结果为FALSE")
//		}
//	}
//}
