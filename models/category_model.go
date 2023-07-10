package models

import "gorm.io/gorm"

type Catagory struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `gorm:"name" json:"name"`
}
