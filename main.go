package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/controllers"
	"github.com/techdenglei/eshop/initiailizers"
)

func init() {
	initiailizers.ConnectToDatabase()
}

func main() {

	router := gin.Default()
	registerRouter(router)
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	router.Run(host + ":" + port)
}

func registerRouter(engine *gin.Engine) {
	new(controllers.CatalogController).Router(engine)
}
