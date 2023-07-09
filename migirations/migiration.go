package main

import (
	"github.com/techdenglei/eshop/initiailizers"
	"github.com/techdenglei/eshop/models"
)

func init() {
	initiailizers.ConnectToDatabase()
}

func main() {
	initiailizers.DB.AutoMigrate(&models.CatalogBrand{}, &models.CatalogItem{}, &models.CatalogType{})
}
