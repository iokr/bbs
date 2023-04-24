package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iokr/bbs/api/rest"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/middleware"
)

func RegisterHTTPServerRouter(router *gin.Engine, indexBiz *biz.IndexBiz,
	userBiz *biz.UserBiz, articleBiz *biz.ArticleBiz) {
	userRest := rest.NewUserRest(userBiz)
	articleRest := rest.NewArticleRest(articleBiz)
	indexRest := rest.NewIndexRest(indexBiz)

	router.Static("/static/", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.LoadHTMLGlob("./api/views/*")

	// index
	router.GET("/", indexRest.IndexPage)
	router.GET("/index", indexRest.IndexLoginPage)

	router.GET("/articles", indexRest.MoreArticles)
	router.GET("/user/:userId/article/:articleId", indexRest.ArticleDetail)

	// user
	router.GET("/login", userRest.LoginPage)
	router.GET("/register", userRest.RegisterPage)

	router.POST("/register", userRest.Register)
	router.POST("/login", userRest.Login)

	// 游客

	// article
	user := router.Group("/user", middleware.CheckUserSession)
	user.GET("/article", articleRest.CreateArticlePage)
	user.POST("/article", articleRest.CreateArticle)
	user.PUT("/article/:id", articleRest.UpdateArticle)
}
