package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/initiailizers"
	"github.com/techdenglei/eshop/routes"
)

func init() {
	initiailizers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	routes.UseRoutes(router)
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	router.Run(host + ":" + port)
}
