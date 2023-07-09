package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ID         int    `json:"id"`
	Content    string `json:"content"`
	QuestionID int    `json:"question_id"`
}
