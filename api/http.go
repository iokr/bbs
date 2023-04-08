package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iokr/bbs/api/rest"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/middleware"
)

func RegisterHTTPServerRouter(router *gin.Engine, userBiz *biz.UserBiz,
	articleBiz *biz.ArticleBiz) {
	userRest := rest.NewUserRest(userBiz)
	articleRest := rest.NewArticleRest(articleBiz)

	// user
	router.POST("/register", userRest.Register)
	router.POST("/login", userRest.Login)

	// 游客

	// article
	user := router.Group("/user", middleware.ParseToken)
	user.POST("/article", articleRest.CreateArticle)
	user.PUT("/article/:id", articleRest.UpdateArticle)
}
