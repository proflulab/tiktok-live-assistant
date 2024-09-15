package services

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"tiktok-live-assistant/configs"
	"tiktok-live-assistant/handlers"
)

// ChromeCtrl 操作浏览器
func ChromeCtrl() {

	// 从env获取URL
	tiktokURL := handlers.GetEnv("TIKTOK_URL")
	liveRoom := handlers.GetEnv("TIKTOK_LIVE_URL")

	var opts []chromedp.ExecAllocatorOption
	opts = configs.BuildChromeDpOpts(opts)

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

	// 启动协程
	//go func() {
	//	// 在协程中运行AskGuard
	//	if err := AskGuard(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	// 执行任务2
	if err := chromedp.Run(ctx, secondTask(liveRoom)); err != nil {
		return
	}

	// 执行任务3
	if err := chromedp.Run(ctx, thirdTask()); err != nil {
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
