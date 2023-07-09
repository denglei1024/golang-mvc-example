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
	c.POST("/question/update", QuestionUpdate)
	c.DELETE("/question/delete/:id", QuestionDelete)
	c.GET("/question/list", QuestionList)
	c.GET("/question/:id", QuestionShow)
}

func QuestionCreate(c *gin.Context) {
	// 从body中获取数据
	reqBody := models.Question{}
	c.Bind(&reqBody)
	question := models.Question{
		ID:          reqBody.ID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		CatalogID:   reqBody.CatalogID,
	}

	// 创建一个question
	result := initiailizers.DB.Create(&question)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, &question)
}

func QuestionUpdate(c *gin.Context) {
	reqBody := models.Question{}
	c.Bind(&reqBody)
	question := models.Question{}
	initiailizers.DB.First(&question, reqBody.ID)
	question.Title = reqBody.Title
	question.Description = reqBody.Description
	question.CatalogID = reqBody.CatalogID
	result := initiailizers.DB.Save(&question)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, &question)
}

func QuestionDelete(c *gin.Context) {
	id := c.Param("id")
	question := models.Question{}
	result := initiailizers.DB.Delete(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.Status(http.StatusOK)
}

func QuestionShow(c *gin.Context) {
	id := c.Param("id")
	question := models.Question{}
	result := initiailizers.DB.First(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, &question)
}

func QuestionList(c *gin.Context) {
	questions := []models.Question{}
	result := initiailizers.DB.Scopes(Paginate(c.Request)).Find(&questions)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, &questions)
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
