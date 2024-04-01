// Package topic 模型
package topic

import (
	"yftxhub/app/models"
	"yftxhub/app/models/category"
	"yftxhub/app/models/user"
	"yftxhub/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title"`
	Body       string `json:"body"`
	UserID     string `json:"user_id"`
	CategoryID string `json:"category_id"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
