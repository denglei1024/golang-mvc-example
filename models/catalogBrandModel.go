package models

import "gorm.io/gorm"

type CatalogBrand struct {
	gorm.Model
	Id    int    `json:"id"`
	Brand string `json:"brand"`
}
