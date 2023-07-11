package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/middlewares"
	"github.com/techdenglei/golang-mvc-example/routes"
)

func init() {
	initiailizers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())
	routes.UseRoutes(router)
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	router.Run(host + ":" + port)
}
