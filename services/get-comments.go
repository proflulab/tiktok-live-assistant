package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

// 存储数据
var dataList []map[string]string

// 存储上一个id
var lastDataID string

// 获取评论
func getComments() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		for {
			// 等待页面元素加载完成
			var htmlContent string
			err = chromedp.Run(ctx,
				chromedp.WaitVisible(`div.webcast-chatroom___item.webcast-chatroom___enter-done`, chromedp.ByQuery),
				chromedp.OuterHTML(`html`, &htmlContent),
			)
			if err != nil {
				log.Fatalf("Failed to get page content: %v", err)
			}

			// 使用goquery解析HTML
			doc, failed := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
			if failed != nil {
				log.Fatalf("Failed to parse HTML: %v", failed)
			}

			doc.Find(`div.webcast-chatroom___item.webcast-chatroom___enter-done`).Each(func(i int, s *goquery.Selection) {
				currentDataID, exists := s.Attr("data-id")
				// 如果 `data-id` 与上一个相同，则跳过
				if !exists || currentDataID == lastDataID {
					return
				}

				// 更新 `last_data_id`
				lastDataID = currentDataID
				// 获取用户名和评论
				username := s.Find(`span.u2QdU6ht`).Text()
				comment := s.Find(`span.WsJsvMP9`).Text()

				// 检查 lEfJhurR 类别并过滤掉
				if s.Find(`span.lEfJhurR`).Length() > 0 {
					return
				}

				// 检查用户名是否以 `：` 结束并处理
				if !strings.HasSuffix(username, "：") {
					return
				}

				username = strings.TrimSuffix(username, "：")

				// 添加数据到 dataList
				dataList = append(dataList, map[string]string{
					"username": username,
					"comment":  comment,
				})

				fmt.Println("=================================")
				fmt.Println("用户名：", username)
				fmt.Println("评论：", comment)

			})

			time.Sleep(time.Second)
		}
	}
}

// 保存 dataList 到项目目录的函数
func saveDataList(dataList []map[string]string) {
	file, err := os.Create("data_list.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 设置 JSON 缩进格式
	if err = encoder.Encode(dataList); err != nil {
		fmt.Printf("Error encoding dataList to JSON: %v\n", err)
	}
}
