package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iokr/bbs/core/middleware"
)

func SetUserSession(c *gin.Context, userId uint) {
	session := &middleware.SessionValue{
		Id: userId,
	}
	middleware.SetLoginCookie(c, session)
}

func CheckUserSession(c *gin.Context) {
	tokenString := middleware.GetCookieSession(c)

	sessionValue, ok := tokenString.Values["Authorization"]
	if !ok {
		c.Abort()
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user := sessionValue.(*middleware.SessionValue)
	has, err := verifyLoginUser(user.Id)
	if err != nil || has == false {
		c.Abort()
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.Set("userId", user.Id)
	c.Next()
}

func verifyLoginEmail(email string) {

}

func verifyLoginUser(userId uint) (bool, error) {
	return true, nil
}
