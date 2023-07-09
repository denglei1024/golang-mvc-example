package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/models"
	"gorm.io/gorm"
)

type QuestionController struct {
}

func (question *QuestionController) Router(c *gin.Engine) {
	c.POST("/question/create", QuestionCreate)
}

func QuestionCreate(c *gin.Context) {
	// 从body中获取数据
	reqBody := models.Question{}
	c.BindJSON(&reqBody)
	question := models.Question{
		Title:       reqBody.Title,
		Description: reqBody.Description,
	}

	// 创建一个question
	result := initiailizers.DB.Create(&question)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, &question)
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
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
