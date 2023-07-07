package models

type CatalogItem struct {
	Id                int          `json:"id"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	Price             float64      `json:"price"`
	PictureFileName   string       `json:"pictureFileName"`
	PictureUri        string       `json:"pictureUri"`
	CatalogTypeId     int          `json:"catalogTypeId"`
	AvailableStock    int          `json:"availableStock"`
	RestockThreshold  int          `json:"restockThreshold"`
	MaxStockThreshold int          `json:"MaxStockThreshold"`
	OnReorder         bool         `json:"onReorder"`
	Brand             CatalogBrand `json:"catalogBrand"`
	Type              CatalogType  `json:"catalogType"`
}

// func (catalogItem *CatalogItem) RemoveStock(quantityDesired int) int {

// }
