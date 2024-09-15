package services

import (
	"fmt"
	"tiktok-live-assistant/handlers"
	"time"
)

var lastPrimaryKey string

func GetAskGuard() {
	res := handlers.GetDataWithMinQuestionTime()
	if res.ID != "" {
		fmt.Println("=================================")
		fmt.Println(res.UserName)
		fmt.Println(res.CommentContent)
		// 如果检查句子通过
		if handlers.SentenceClassify(res.CommentContent) {
			// 将question_judgment 更新为true
			handlers.UpdateQuestionJudgment(res.ID, true)
		} else {
			handlers.UpdateQuestionJudgment(res.ID, false)
		}
	} else {
		fmt.Println("No records found with question_judgment as NULL.")
		time.Sleep(2 * time.Second)
	}
}
