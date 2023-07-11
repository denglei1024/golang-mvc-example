package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/techdenglei/golang-mvc-example/initiailizers"
	"github.com/techdenglei/golang-mvc-example/models"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// get route path off the request
		path := ctx.Request.URL.Path

		// ignore it if the path contains "login" and "signup"
		if path == "/login" || path == "/signup" {
			ctx.Next()
			return
		}
		// get the token from the authorization
		authorization := ctx.Request.Header.Get("Authorization")

		// check if the authorization is empty
		if authorization == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		authHeader := strings.Split(authorization, "Bearer ")
		if len(authHeader) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := authHeader[1]
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// check if the exp off claims are valid
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}
			// get user with token sub
			var user models.User
			initiailizers.DB.First(&user, claims["sub"])
			if user.ID == 0 {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}
			ctx.Set("user", user)
			ctx.Next()

		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
