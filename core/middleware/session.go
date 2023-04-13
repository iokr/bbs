package middleware

import (
	"encoding/gob"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type SessionValue struct {
	Id uint
}

var Store = sessions.NewCookieStore([]byte("abc!@#$%^&*1234567890"))

func init() {
	gob.Register(&SessionValue{})
}

func SetLoginCookie(c *gin.Context, sessionValue *SessionValue) {
	Store.Options.HttpOnly = true

	Store.MaxAge(10000) // 秒

	session := GetCookieSession(c)
	session.Values["Authorization"] = sessionValue
	session.Save(c.Request, c.Writer)
}

func GetCookieSession(c *gin.Context) *sessions.Session {
	session, _ := Store.Get(c.Request, "bbs_token")
	return session
}
