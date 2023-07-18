package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/models"
	"github.com/techdenglei/golang-mvc-example/utils"
	"gorm.io/gorm"
)

type QuestionController struct {
}

func (question *QuestionController) Router(c *gin.Engine) {
	c.POST("/question/create", questionCreate)
	c.POST("/question/update", questionUpdate)
	c.DELETE("/question/delete/:id", questionDelete)
	c.GET("/question/list", questionList)
	c.GET("/question/:id", questionShow)
}

func questionCreate(c *gin.Context) {
	// get data off the request body
	reqBody := models.Question{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.Error(c, err)
		return
	}
	question := models.Question{
		ID:          reqBody.ID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		CategoryID:  reqBody.CategoryID,
	}

	// create a new question
	result := initiailizers.DB.Create(&question)
	if result.Error != nil {
		utils.Error(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, &question)
}

func questionUpdate(c *gin.Context) {
	reqBody := models.Question{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.Error(c, err)
		return
	}
	question := models.Question{}
	initiailizers.DB.First(&question, reqBody.ID)
	question.Title = reqBody.Title
	question.Description = reqBody.Description
	question.CategoryID = reqBody.CategoryID
	result := initiailizers.DB.Save(&question)
	if result.Error != nil {
		utils.Error(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, &question)
}

func questionDelete(c *gin.Context) {
	id := c.Param("id")
	question := models.Question{}
	result := initiailizers.DB.Delete(&question, id)
	if result.Error != nil {
		utils.Error(c, result.Error)
		return
	}
	c.Status(http.StatusOK)
}

func questionShow(c *gin.Context) {
	id := c.Param("id")
	question := models.Question{}
	result := initiailizers.DB.First(&question, id)
	if result.Error != nil {
		utils.Error(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, &question)
}

func questionList(c *gin.Context) {
	questions := []models.Question{}
	result := initiailizers.DB.Scopes(paginate(c.Request)).Find(&questions)
	if result.Error != nil {
		utils.Error(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, &questions)
}

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
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
