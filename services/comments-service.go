package services

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"tiktok-live-assistant/handlers"
	"time"
)

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
			username = strings.TrimSuffix(username, "：")
			// 添加数据到数据库
			uniqueId := handlers.GenerateRandomString(32) // 获取唯一表示
			now := time.Now()                             // 获取时间戳
			handlers.InsertData(uniqueId, username, comment, now)

			GetAskGuard()

			//fmt.Println("=================================")
			//fmt.Println("用户名：", username)
			//fmt.Println("评论：", comment)
			// 等待 1 秒钟再进行下一次抓取
			time.Sleep(1 * time.Second)
		}
	}
}
