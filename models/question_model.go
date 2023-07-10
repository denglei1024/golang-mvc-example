package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}
