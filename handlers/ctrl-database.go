package handlers

import (
	"tiktok-live-assistant/database"
	"tiktok-live-assistant/models"
)

// InsertData 新增数据
func InsertData(id, username, questionTime, commentContent string) {
	product := models.Product{ID: id, UserName: username, QuestionTime: questionTime, CommentContent: commentContent}
	database.DB.Create(&product)
}

// DeleteData 删除数据
func DeleteData(id string) {
	database.DB.Delete(&models.Product{ID: id})
}

// UpdateDate 修改数据
func UpdateDate(id string, questionJudgment bool) {
	var product models.Product
	database.DB.First(&product, models.Product{ID: id})
	database.DB.Model(&product).Update("QuestionJudgment", questionJudgment)
}

// GetDataByID 查询(通过ID)
func GetDataByID(id string) models.Product {
	var product models.Product
	database.DB.First(&product, models.Product{ID: id})
	return product
}
