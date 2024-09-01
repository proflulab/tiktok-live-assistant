package services

import (
	"context"
	"github.com/chromedp/cdproto/network"
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
		chromedp.UserDataDir("D:\\.cache\\chrome-dp\\Directory"),
		chromedp.Flag("disk-cache-dir", "D:\\.cache\\chrome-dp\\UserCache"),
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"),
		chromedp.Flag("lang", "en-US,en;q=0.9"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("timezone", "America/New_York"),
	)

	// 创建一个分配器上下文，用于管理Chrome实例的执行
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 建一个新上下文，用于执行浏览器任务
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 执行任务1
	if err := chromedp.Run(ctx, firstTask(tiktokURL)); err != nil {
		log.Fatal(err)
		return
	}

	// 执行任务2
	if err := chromedp.Run(ctx, secondTask(liveRoom)); err != nil {
		return
	}

	// 清空所有 Cookies
	e := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		return network.ClearBrowserCookies().Do(ctx)
	}))

	if e != nil {
		log.Fatalf("Failed to clear cookies: %v", e)
	}

}
