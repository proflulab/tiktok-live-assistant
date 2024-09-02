package configs

import "github.com/chromedp/chromedp"

func BuildChromeDpOpts(opts []chromedp.ExecAllocatorOption) []chromedp.ExecAllocatorOption {
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
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
	return opts
}
