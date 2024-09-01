package services

import (
	"github.com/chromedp/chromedp"
	"time"
)

// 第一个任务
func firstTask(url string) chromedp.Tasks {
	return chromedp.Tasks{

		// 加载cookies (如果有)
		loadCookies(),

		// 打开网站
		chromedp.Navigate(url),

		// 自动操作：选择账号密码登录的Tab
		//chromedp.Click(`#web-login-container > article > article > article > div > ul.web-login-tab-list > li:nth-child(3)`),

		// 是否保存Cookies
		checkSaveCookies(),
	}
}

func secondTask(live string) chromedp.Tasks {
	return chromedp.Tasks{
		// 进入直播间
		chromedp.Navigate(live),

		// 等待页面加载
		chromedp.Sleep(3 * time.Second),

		getComments(),
	}
}
