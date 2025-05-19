package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
      var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func HashPassword(password string) string {

}
func Verifypassword(userPassword string, providedPassword string) (bool, string) {

}
