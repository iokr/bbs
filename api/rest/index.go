package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/response"
)

type IndexRest struct {
	indexBiz *biz.IndexBiz
}

func NewIndexRest(indexBiz *biz.IndexBiz) *IndexRest {
	return &IndexRest{
		indexBiz: indexBiz,
	}
}

func (r *IndexRest) IndexPage(c *gin.Context) {
	respMap := make(map[string]interface{})
	respMap["site_name"] = "BBS"

	c.HTML(http.StatusOK, "index.html", respMap)
}

func (r *IndexRest) MoreArticles(c *gin.Context) {
	var param request.MoreArticleRequest

	err := c.MustBindWith(&param, binding.Query)
	if err != nil {
		response.ServerJson(c.Writer, nil, err)
		return
	}

	// 每次最多只查询10条数据.
	param.Limit = 10
	result, err := r.indexBiz.FindMoreArticles(c, &param)
	if err != nil {
		response.ServerJson(c.Writer, nil, err)
		return
	}
	response.ServerJson(c.Writer, result, nil)
}

func (r *IndexRest) ArticleDetail(c *gin.Context) {
	respMap := make(map[string]interface{})
	respMap["site_name"] = "BBS"

	c.HTML(http.StatusOK, "article_detail.html", respMap)
}
