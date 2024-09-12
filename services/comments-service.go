package services

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"tiktok-live-assistant/handlers"
	"time"
	"unicode/utf16"
)

// 全局变量，用于存储评论数据
var dataList [][]string

// 存储上一个id
var lastDataID string

// 获取评论
func getComments() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		for {
			var htmlContent string
			var doc *goquery.Document
			var selection *goquery.Selection
			var currentDataID string
			// 获取页面
			htmlContent, ctx, err = handlers.GetHttpHtmlContent(htmlContent, ctx)
			if err != nil {
				log.Fatalf("Failed to get page content: %v", err)
			}
			// 解析数据
			doc, err = handlers.GetSpecialData(htmlContent, doc)
			if err != nil {
				log.Fatalf("Failed to parse HTML: %v", err)
			}
			selection = handlers.GetCurrentSelection(selection, doc)
			currentDataID, _ = selection.Attr("data-id")
			if currentDataID == lastDataID {
				continue
			}
			lastDataID = currentDataID
			// 检查 lEfJhurR 类别并过滤掉
			if selection.Find(`span.lEfJhurR`).Length() > 0 {
				continue
			}
			// 获取用户名和评论
			username := selection.Find(`span.u2QdU6ht`).Text()
			comment := selection.Find(`span.WsJsvMP9`).Text()
			// 检查用户名是否以 `：` 结束并处理
			if !strings.HasSuffix(username, "：") {
				continue
			}
			// 去掉用户名中的最后一个字符 `：`
			username = strings.TrimSuffix(username, "：")

			// 这段代码用于判断是否是机器人发送的回复，防止信息被录入
			if IsRobotReply(comment, 4) {
				continue // 跳过该条记录
			}
			// 添加数据到切片
			// 将 username 和 comment 添加到 dataList
			dataList = append(dataList, []string{username, comment, "", ""})

			// 添加数据到数据库
			uniqueId := handlers.GenerateRandomString(32) // 获取唯一表示
			now := time.Now()                             // 获取时间戳
			handlers.InsertData(uniqueId, username, comment, now)
			// 检查句子
			GetAskGuard()

			// 等待 1 秒钟再进行下一次抓取
			time.Sleep(1 * time.Second)
		}
	}
}

// 删除非 BMP 字符的函数
func removeNonBMPCharacters(text string) string {
	var result []rune
	for _, c := range text {
		if utf16.IsSurrogate(c) || c > 0xFFFF {
			continue
		}
		result = append(result, c)
	}
	return string(result)
}

// IsRobotReply 判断是否是机器人回复的函数
func IsRobotReply(comment string, maxCheckCount int) bool {
	if len(dataList) > 1 {
		var latestNonEmptyComment string
		// 从下往上遍历 dataList，最多检查 maxCheckCount 行
		for i := len(dataList) - 1; i >= 0 && i >= len(dataList)-maxCheckCount; i-- {
			row := dataList[i]
			if row[3] != "" { // 如果当前行的第4项不为空
				latestNonEmptyComment = removeNonBMPCharacters(row[3])
				//fmt.Println("Checking row: ", i+1)
				//fmt.Println("Latest Non-Empty Comment: ", latestNonEmptyComment)
				break // 找到第一个非空值项后跳出循环
			}
		}
		// 如果找到了非空值项，则与 comment 的前10个字符进行比较
		if latestNonEmptyComment != "" && len(comment) >= 10 {
			//fmt.Println("Comparing: ", latestNonEmptyComment[:10])
			//fmt.Println("Comment: ", comment[:10])
			if comment[:10] == latestNonEmptyComment[:10] {
				//fmt.Println("Latest row number: ", len(dataList))
				return true // 是机器人发送的回复
			}
		}
	}

	return false // 不是机器人发送的回复
}
