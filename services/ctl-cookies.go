package services

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var cookiesPath string

func init() {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 将相对路径与当前工作目录结合，生成绝对路径
	cookiesPath = filepath.Join(cwd, "public", "cookies", "cookies.tmp")
}

// 加载Cookies
func loadCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat(cookiesPath); os.IsNotExist(_err) {
			return
		}
		// 如果存在则读取cookies的数据
		cookiesData, err := ioutil.ReadFile(cookiesPath)
		if err != nil {
			return
		}
		// 反序列化
		cookiesParams := network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
			return
		}

		// 设置cookies
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

// 保存Cookies
func saveCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {

		fmt.Println("开始保存......")
		// cookies的获取对应是在devTools的network面板中
		// 1. 获取cookies
		cookies, err := network.GetCookies().Do(ctx)
		if err != nil {
			return
		}

		// 2. 序列化
		cookiesData, err := network.GetCookiesReturns{Cookies: cookies}.MarshalJSON()
		if err != nil {
			return
		}

		// 3. 存储到临时文件
		if err = ioutil.WriteFile(cookiesPath, cookiesData, 0755); err != nil {
			return
		}
		fmt.Println("已保存")
		return
	}
}

// 是否保存Cookies
func checkSaveCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var choice int
		fmt.Println("请完成登录后再来选择")
		fmt.Println("是否保存Cookies？ 1.是  2.否")
		fmt.Println("请输入数字：")
		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("发生错误", err.Error())
			return
		}
		if choice == 1 {
			saveCookies()
		} else if choice == 2 {
			return
		} else {
			fmt.Println("格式错误")
			return
		}
		return
	}
}
