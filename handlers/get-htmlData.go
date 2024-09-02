package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"os"
	"strings"
)

// RoomXPATH 元素
var RoomXPATH = GetRoomXPATH()

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

// SaveDataList 保存 dataList 到项目目录的函数
func SaveDataList(dataList []map[string]string) {
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
