package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/initiailizers"
	"github.com/techdenglei/eshop/models"
)

type CatalogController struct {
}

func (catalog *CatalogController) Router(c *gin.Engine) {
	c.GET("/catalogtypes", catalogTypes)
	c.GET("/catalogbrands", catalogBrands)
	c.GET("/items", items)
	c.GET("/items/:id", itemById)
	c.POST("/items", createProduct)
}

func catalogTypes(c *gin.Context) {
	c.String(http.StatusOK, "from catalogtypes")
}

func catalogBrands(c *gin.Context) {
	c.String(http.StatusOK, "from catalogbrands")
}

func items(c *gin.Context) {
	c.String(http.StatusOK, "from items")
}

func itemById(c *gin.Context) {
	c.String(http.StatusOK, "from item by id")
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
