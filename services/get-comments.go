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

// RoomXPATH 元素
var RoomXPATH = GetRoomXPATH()

// 获取评论
func getComments() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		for {
			var htmlContent string
			var doc *goquery.Document
			var selection *goquery.Selection
			var currentDataID string
			// 获取页面
			htmlContent, ctx, err = GetHttpHtmlContent(htmlContent, ctx)
			if err != nil {
				log.Fatalf("Failed to get page content: %v", err)
			}
			// 解析数据
			doc, err = GetSpecialData(htmlContent, doc)
			if err != nil {
				log.Fatalf("Failed to parse HTML: %v", err)
			}
			selection = GetCurrentSelection(selection, doc)
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
			// 添加数据到 dataList
			dataList = append(dataList, map[string]string{
				"username": username,
				"comment":  comment,
			})
			fmt.Println("=================================")
			fmt.Println("用户名：", username)
			fmt.Println("评论：", comment)
			// 等待 1 秒钟再进行下一次抓取
			time.Sleep(1 * time.Second)
		}
	}
}

// GetHttpHtmlContent 抓取网站动态数据
func GetHttpHtmlContent(htmlContent string, ctx context.Context) (string, context.Context, error) {
	// 等待页面元素加载完成
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(RoomXPATH, chromedp.ByQuery),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	)
	return htmlContent, ctx, err
}

// GetSpecialData 解析为具体数据
func GetSpecialData(htmlContent string, doc *goquery.Document) (*goquery.Document, error) {
	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	return doc, err
}

// GetCurrentSelection 获取当前选择器
func GetCurrentSelection(selection *goquery.Selection, doc *goquery.Document) *goquery.Selection {
	doc.Find(RoomXPATH).Each(func(i int, s *goquery.Selection) {
		selection = s
	})
	return selection
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
