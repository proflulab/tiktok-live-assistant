package services

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

// ChromeCtrl 操作浏览器
func ChromeCtrl() {

	// 从env获取URL
	tiktokURL := GetTiktokURL()
	liveRoom := GetTiktokLiveURL()

	// 使用chrome-dp 配置Chrome浏览器的选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
	)

	// 创建一个分配器上下文，用于管理Chrome实例的执行
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	// 建一个新上下文，用于执行浏览器任务
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)

	if err := chromedp.Run(ctx, firstTask(tiktokURL, liveRoom)); err != nil {
		log.Fatal(err)
		return
	}
}
