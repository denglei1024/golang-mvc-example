package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatalogController struct {
}

func (catalog *CatalogController) Router(c *gin.Engine) {
	c.GET("/catalogtypes", catalogTypes)
	c.GET("/catalogbrands", catalogBrands)
	c.GET("/items", items)
	c.GET("/items/:id", itemById)
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
