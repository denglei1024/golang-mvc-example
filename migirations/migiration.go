package main

import (
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/models"
)

func init() {
	initiailizers.ConnectToDatabase()
}

func main() {
	initiailizers.DB.AutoMigrate(&models.Question{}, &models.Answer{}, &models.Catagory{})
	initiailizers.DB.AutoMigrate(&models.User{})
}
