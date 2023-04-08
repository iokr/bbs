package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iokr/bbs/core/middleware"
)

const (
	secretKey = "asd!@#$%^&*123456789"
)

func GenerateToken(claims map[string]interface{}) (token string, err error) {
	opts := []middleware.JWTFilterOptions{
		middleware.WithSecretKey([]byte(secretKey)),
	}
	jwtFilter := middleware.NewJWTFilter(opts...)
	return jwtFilter.GenerateToken(claims)
}

func ParseToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	opts := []middleware.JWTFilterOptions{
		middleware.WithSecretKey([]byte(secretKey)),
	}
	jwtFilter := middleware.NewJWTFilter(opts...)
	claims, err := jwtFilter.ParseToken(tokenString)
	if err != nil {
		c.Abort()
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}

	userId := uint(claims["userId"].(float64))
	// todo 验证userId是否正确
	c.Set("userId", userId)
	c.Next()
}
