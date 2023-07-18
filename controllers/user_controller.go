package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/models"
	"github.com/techdenglei/golang-mvc-example/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

func (c *UserController) Router(engine *gin.Engine) {
	engine.POST("/signup", signup)
	engine.POST("/login", login)
}

func signup(c *gin.Context) {

	// get data off the request body
	var userReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userReq); err != nil {
		utils.Error(c, err)
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		utils.Error(c, err)
		return
	}

	// create a new user
	newUser := models.User{
		Email:    userReq.Email,
		Password: string(hashedPassword),
	}

	// save the new user
	initiailizers.DB.Create(&newUser)
	utils.Success(c, nil)
}

// define a method for login
func login(c *gin.Context) {
	// get data off the request body
	var userReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userReq); err != nil {
		utils.Error(c, err)
		return
	}

	// look up the requested user
	var user models.User
	initiailizers.DB.Where("email =?", userReq.Email).First(&user)

	if user.ID == 0 {
		utils.Fail(c, 404, "User not found")
		return
	}

	// compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		utils.Fail(c, 401, "Invalid password")
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		utils.Error(c, err)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	// set cookie
	c.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)
	utils.Success(c, nil)
}
