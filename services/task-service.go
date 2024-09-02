package services

import (
	"github.com/chromedp/chromedp"
	"tiktok-live-assistant/handlers"
	"time"
)

// 第一个任务
func firstTask(url string) chromedp.Tasks {
	return chromedp.Tasks{

		// 加载cookies (如果有)
		handlers.LoadCookies(),

		// 打开网站
		chromedp.Navigate(url),

		// 自动操作：选择账号密码登录的Tab
		//chromedp.Click(`#web-login-container > article > article > article > div > ul.web-login-tab-list > li:nth-child(3)`),

		// 是否保存Cookies
		handlers.CheckSaveCookies(),
	}
}

func secondTask(live string) chromedp.Tasks {
	return chromedp.Tasks{
		// 进入直播间
		chromedp.Navigate(live),

		// 等待页面加载
		chromedp.Sleep(2 * time.Second),

		getComments(),
	}
}
