package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/techdenglei/golang-mvc-example/controllers"
)

func UseRoutes(engine *gin.Engine) {
	new(controllers.QuestionController).Router(engine)
	// register user routes
	new(controllers.UserController).Router(engine)
}
