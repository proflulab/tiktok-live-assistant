package models

import "time"

type Product struct {
	ID               string    `gorm:"primaryKey"`
	UserName         string    // UserName 用户名
	QuestionTime     time.Time // QuestionTime 提问的时间戳
	CommentContent   string    // CommentContent 评论内容
	QuestionJudgment bool      // QuestionJudgment 是否是问题
	AnswerContent    string    // AnswerContent 回答内容
}
