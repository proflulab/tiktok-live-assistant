package handlers

import (
	"tiktok-live-assistant/configs"
	"tiktok-live-assistant/models"
	"time"
)

// InsertData 新增数据
func InsertData(id, username, commentContent string, questionTime time.Time) {
	product := models.Product{ID: id, UserName: username, QuestionTime: questionTime, CommentContent: commentContent}
	configs.DB.Create(&product)
}

// DeleteData 删除数据
func DeleteData(id string) {
	configs.DB.Delete(&models.Product{ID: id})
}

// UpdateQuestionJudgment 修改数据
func UpdateQuestionJudgment(id string, questionJudgment bool) {
	var product models.Product
	configs.DB.First(&product, models.Product{ID: id})
	configs.DB.Model(&product).Update("QuestionJudgment", questionJudgment)
}

// GetDataByID 查询(通过ID)
func GetDataByID(id string) models.Product {
	var product models.Product
	configs.DB.First(&product, models.Product{ID: id})
	return product
}

func GetDataWithMinQuestionTime() models.Product {
	var product models.Product
	// 查询 question_judgment 为 false 或 0，按 question_time 升序排序，获取第一条记录
	configs.DB.Where("question_judgment = 0 OR question_judgment IS NULL", false).
		Order("question_time DESC").
		Limit(1).
		First(&product)
	return product
}

func GetUserComments() []models.Product {
	var products []models.Product
	configs.DB.Select("user_name", "comment_content").
		Find(&products)
	return products
}
