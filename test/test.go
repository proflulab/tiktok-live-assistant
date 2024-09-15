package test

import (
	"fmt"
	"tiktok-live-assistant/handlers"
	"tiktok-live-assistant/services"
)

// Unit01 测试句子分类能否正常工作
func Unit01() {
	for {
		fmt.Println("请输入句子，或输入退出结束测试")
		var sentence string
		_, err := fmt.Scanln(&sentence)
		if err != nil {
			return
		}
		if sentence == "退出" {
			break
		}
		flag := handlers.SentenceClassify(sentence)
		if flag {
			fmt.Println("结果为TRUE")
		} else {
			fmt.Println("结果为FALSE")
		}
	}
}

// Unit02 测试句子是否是机器人回复的
func Unit02() {
	var dataList [][]string
	// 获取数据
	products := handlers.GetUserComments()

	// 读取数据
	for _, product := range products {
		dataList = append(dataList, []string{product.UserName, product.CommentContent})
	}

	// 遍历 dataList，提取用户名和评论，执行 isRobotReply 判断
	for _, data := range dataList {
		username := data[0]
		comment := data[1]
		isRobot := services.IsRobotReply(comment, 4)

		fmt.Printf("User: %s, Comment: %s, Is robot reply: %v\n", username, comment, isRobot)
	}
}

// Unit03 测试调用大模型
func Unit03(question string) {
	resp := services.SendAndGetReply(question)
	for _, content := range resp {
		fmt.Println(content)
	}
}
