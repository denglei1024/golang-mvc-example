package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/initiailizers"
	"github.com/techdenglei/eshop/models"
	"gorm.io/gorm"
)

type CatalogController struct {
}

func (catalog *CatalogController) Router(c *gin.Engine) {
	c.GET("/catalogtypes", catalogTypes)
	c.GET("/catalogbrands", catalogBrands)
	c.GET("/items", items)
	c.GET("/items/:id", itemById)
	c.POST("/items", createProduct)
	c.DELETE("/items/:id", deleteProduct)
}

func catalogTypes(c *gin.Context) {
	result := initiailizers.DB.Find(&models.CatalogType{})
	c.JSON(http.StatusOK, result)
}

func catalogBrands(c *gin.Context) {
	result := initiailizers.DB.Find(&models.CatalogBrand{})
	c.JSON(http.StatusOK, result)
}

func items(c *gin.Context) {
	products := []models.CatalogItem{}
	initiailizers.DB.Scopes(Paginate(c.Request)).Find(&products)
	c.JSON(http.StatusOK, &products)
}

func itemById(c *gin.Context) {
	id := c.Param("id")
	catalogItem := []models.CatalogItem{}
	initiailizers.DB.First(&catalogItem, id)
	c.JSON(http.StatusOK, &catalogItem[0])
}

func createProduct(c *gin.Context) {
	var product models.CatalogItem
	c.BindJSON(&product)
	result := initiailizers.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	initiailizers.DB.Delete(&models.CatalogItem{}, id)
	c.Status(http.StatusOK)
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
