package services

import "github.com/chromedp/chromedp"

// 第一个任务
func firstTask(url, live string) chromedp.Tasks {
	return chromedp.Tasks{

		// 加载cookies (如果有)
		loadCookies(),

		// 打开网站
		chromedp.Navigate(url),

		// 是否保存Cookies
		checkSaveCookies(),

		// 进入直播间
		chromedp.Navigate(live),

		// 选择账号密码登录的Tab
		//chromedp.Click(`#web-login-container > article > article > article > div > ul.web-login-tab-list > li:nth-child(3)`),

	}
}
