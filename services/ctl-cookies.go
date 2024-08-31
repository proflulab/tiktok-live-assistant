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
		fmt.Println("正在读取Cookies...")
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat(cookiesPath); os.IsNotExist(_err) {
			fmt.Println("未找到Cookies")
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
		//fmt.Println("反序列化后的 Cookies:")
		//cookiesJSON, _ := json.MarshalIndent(cookiesParams.Cookies, "", "  ")
		//fmt.Println(string(cookiesJSON))
		fmt.Println("读取完成，开始设置Cookies")
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

// 是否保存Cookies
func checkSaveCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var choice int
		fmt.Println("完成登录后按任意键继续...")
		_, _ = fmt.Scanln()
		fmt.Println("是否保存Cookies？（如果之前已经保存，可以选否） ")
		fmt.Println("请选择然后按回车键 1.是  2.否")
		_, _ = fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("开始保存......")
			// cookies的获取对应是在devTools的network面板中
			// 1. 获取cookies
			cookies, res := network.GetCookies().Do(ctx)
			if res != nil {
				return
			}
			// 2. 序列化
			cookiesData, res := network.GetCookiesReturns{Cookies: cookies}.MarshalJSON()
			if res != nil {
				return
			}
			// 3. 存储到临时文件
			if res = ioutil.WriteFile(cookiesPath, cookiesData, 0755); res != nil {
				return
			}
			fmt.Println("已保存")
			return
		} else if choice == 2 {
			return
		} else {
			fmt.Println("格式错误")
			return
		}
	}
}
