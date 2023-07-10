package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (c *UserController) Router(engine *gin.Engine) {
	engine.POST("/signup", signup)
}

func signup(c *gin.Context) {

}
