package models

import "gorm.io/gorm"

type CatalogItem struct {
	gorm.Model
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	PictureFileName   string  `json:"pictureFileName"`
	PictureUri        string  `json:"pictureUri"`
	CatalogTypeId     int     `json:"catalogTypeId"`
	AvailableStock    int     `json:"availableStock"`
	RestockThreshold  int     `json:"restockThreshold"`
	MaxStockThreshold int     `json:"MaxStockThreshold"`
	OnReorder         bool    `json:"onReorder"`
	BrandId           int     `json:"brandId"`
	TypeId            int     `json:"typeId"`
}

// func (catalogItem *CatalogItem) RemoveStock(quantityDesired int) int {

// }
