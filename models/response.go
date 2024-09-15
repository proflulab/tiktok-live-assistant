package models

import "go/types"

// Response 响应结构体如下
type Response struct {
	Data ChatResponse `json:"data"`
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
}

type ChatResponse struct {
	Id             string    `json:"id"`
	ConversationId string    `json:"conversation_id"`
	BotId          string    `json:"bot_id"`
	CreateAt       int64     `json:"create_at"`
	CompletedAt    int64     `json:"completed_at"`
	LastError      types.Nil `json:"last_error"`
	MetaData       struct{}  `json:"meta_data"`
	Status         string    `json:"status"`
	Usage          struct {
		TokenCount  int `json:"token_count"`
		OutputCount int `json:"output_count"`
		InputCount  int `json:"input_count"`
	} `json:"usage"`
}

type Messages struct {
	BotId          string `json:"bot_id"`
	Content        string `json:"content"`
	ContentType    string `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Id             string `json:"id"`
	Role           string `json:"role"`
	Type           string `json:"type"`
}
type MessagesResp struct {
	Data []Messages `json:"data"`
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
}

type AdditionalMessages struct {
	Role        string `json:"role"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}

type Payload struct {
	BotId              string               `json:"bot_id"`
	UserId             string               `json:"user_id"`
	Stream             bool                 `json:"stream"`
	AutoSaveHistory    bool                 `json:"auto_save_history"`
	AdditionalMessages []AdditionalMessages `json:"additional_messages"`
}
