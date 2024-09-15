package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"golang.org/x/net/context"
	"tiktok-live-assistant/models"
)

var BASEURL = GetEnv("BASEURL")
var BOT = GetEnv("BOT")
var AUTH = GetEnv("AUTH")

const USERID = "123456789"

func buildPayload(message string) string {
	additionalMessages := models.AdditionalMessages{
		Role:        "user",
		Content:     message,
		ContentType: "text",
	}
	additionalMessageList := []models.AdditionalMessages{additionalMessages}

	payload := models.Payload{
		BotId:              BOT,
		UserId:             USERID,
		Stream:             false,
		AutoSaveHistory:    true,
		AdditionalMessages: additionalMessageList,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	return string(jsonPayload)
}

func SendMsg(message string) models.Response {
	payload := buildPayload(message)
	c, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()

	// Client Request 配置

	req.SetRequestURI(BASEURL)
	req.SetMethod("POST")
	req.SetHeader("Authorization", AUTH)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody([]byte(payload))

	_ = c.Do(context.Background(), req, resp)

	// 解析响应体
	var response models.Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func GetRetrieve(chatId, conversationId string) models.Response {
	url := BASEURL + "/retrieve?conversation_id=" + conversationId + "&chat_id=" + chatId
	c, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(url)
	req.SetMethod("GET")
	req.SetHeader("Authorization", AUTH)
	req.SetHeader("Content-Type", "application/json")

	_ = c.Do(context.Background(), req, resp)
	// 解析响应体
	var response models.Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		fmt.Println(err)
	}

	return response
}

func GetMessageList(chatId, conversationId string) models.MessagesResp {
	url := BASEURL + "/message/list?conversation_id=" + conversationId + "&chat_id=" + chatId
	c, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(url)
	req.SetMethod("GET")
	req.SetHeader("Authorization", AUTH)
	req.SetHeader("Content-Type", "application/json")

	_ = c.Do(context.Background(), req, resp)
	// 解析响应体
	var response models.MessagesResp
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
