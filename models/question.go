package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CatalogID   uint   `json:"catalog_id"`
}
