package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/errors"
	"github.com/iokr/bbs/internal/response"
	"github.com/spf13/cast"
)

type ArticleRest struct {
	articleBiz *biz.ArticleBiz
}

func NewArticleRest(articleBiz *biz.ArticleBiz) *ArticleRest {
	return &ArticleRest{articleBiz: articleBiz}
}

func (r *ArticleRest) CreateArticlePage(c *gin.Context) {
	respMap := make(map[string]interface{})
	c.HTML(http.StatusOK, "article_add.html", respMap)
}

func (r *ArticleRest) CreateArticle(ctx *gin.Context) {
	var param request.CreateArticleRequest

	err := ctx.MustBindWith(&param, binding.Form)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}

	userId, isExist := ctx.Get("userId")
	if !isExist || userId == nil {
		response.ServerJson(ctx.Writer, nil, errors.UserNotFound)
		return
	}

	param.UserId = userId.(uint)
	err = r.articleBiz.CreateArticle(ctx, &param)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}
	response.ServerJson(ctx.Writer, nil, nil)
}

func (r *ArticleRest) UpdateArticle(ctx *gin.Context) {
	var param request.UpdateArticleRequest

	param.Id = cast.ToUint(ctx.Param("id"))
	if param.Id <= 0 {
		response.ServerJson(ctx.Writer, nil, errors.ArticleNotFound)
		return
	}

	err := ctx.MustBindWith(&param, binding.Form)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}

	userId, isExist := ctx.Get("userId")
	if !isExist || userId == nil {
		response.ServerJson(ctx.Writer, nil, errors.UserNotFound)
		return
	}

	param.UserId = cast.ToUint(userId)
	err = r.articleBiz.UpdateArticle(ctx, &param)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}
	response.ServerJson(ctx.Writer, nil, nil)
}
