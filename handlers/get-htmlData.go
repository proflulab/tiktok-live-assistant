package handlers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"strings"
)

// RoomXPATH 元素
var RoomXPATH = GetEnv("ROOM_XPATH")

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
