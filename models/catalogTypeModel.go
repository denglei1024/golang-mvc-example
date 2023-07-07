package models

import "gorm.io/gorm"

type CatalogType struct {
	gorm.Model
	Id   int    `json:"id"`
	Type string `json:"type"`
}
