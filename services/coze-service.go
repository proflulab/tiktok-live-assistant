package services

import (
	"fmt"
	"tiktok-live-assistant/handlers"
	"time"
)

// SendAndGetReply /*
func SendAndGetReply(message string) []string {
	resp := handlers.SendMsg(message)
	if resp.Code != 0 {
		fmt.Println("Error sending message:" + resp.Msg)
	}
	//fmt.Println(resp)
	chatId := resp.Data.Id
	conversationId := resp.Data.ConversationId
	//fmt.Println(chatId, conversationId)

	for {
		messageStatus := handlers.GetRetrieve(chatId, conversationId)
		if messageStatus.Code != 0 {
			fmt.Println("Error sending message: " + messageStatus.Msg)
		}
		status := messageStatus.Data.Status
		//fmt.Println(status)
		if status == "completed" {
			break
		} else {
			// 延时 0.2秒
			time.Sleep(2 * 100 * time.Millisecond)
		}
	}

	messagesResp := handlers.GetMessageList(chatId, conversationId)
	//fmt.Println(messagesResp)
	if messagesResp.Code == 0 && messagesResp.Data != nil {
		// 使用 for 循环和条件语句来筛选 type 为 "answer" 的消息
		var answerContents []string
		for _, msg := range messagesResp.Data {
			if msg.Type == "answer" {
				answerContents = append(answerContents, msg.Content)
			}
		}
		return answerContents
	} else {
		fmt.Println("Error sending message: " + messagesResp.Msg)
	}
	return nil
}

/**
@describe: 回复指定评论
@param: 用户名
@param: 回答
*/
