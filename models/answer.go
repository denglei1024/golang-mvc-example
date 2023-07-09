package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	QuestionID uint   `json:"question_id"`
}
