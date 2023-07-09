package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/controllers"
)

func UseRoutes(engine *gin.Engine) {
	new(controllers.CatalogController).Router(engine)
}
