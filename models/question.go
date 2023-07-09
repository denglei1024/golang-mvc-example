package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CatalogID   int    `json:"catalog_id"`
}
