package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/middleware"
	"github.com/iokr/bbs/internal/response"
)

type UserRest struct {
	userBiz *biz.UserBiz
}

func NewUserRest(userBiz *biz.UserBiz) *UserRest {
	return &UserRest{userBiz: userBiz}
}

func (r *UserRest) LoginPage(c *gin.Context) {
	respMap := make(map[string]interface{})
	c.HTML(http.StatusOK, "login.html", respMap)
}

func (r *UserRest) RegisterPage(c *gin.Context) {
	respMap := make(map[string]interface{})
	c.HTML(http.StatusOK, "register.html", respMap)
}

func (r *UserRest) Register(ctx *gin.Context) {
	var param request.RegisterRequest

	err := ctx.MustBindWith(&param, binding.Form)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}
	response.ServerJson(ctx.Writer, nil, r.userBiz.Register(ctx, &param))
}

func (r *UserRest) Login(c *gin.Context) {
	var param request.LoginRequest

	err := c.MustBindWith(&param, binding.Form)
	if err != nil {
		response.ServerJson(c.Writer, nil, err)
		return
	}

	result, err := r.userBiz.Login(c, &param)
	if err != nil {
		response.ServerJson(c.Writer, nil, err)
		return
	}

	middleware.SetUserSession(c, result.UserId)

	response.ServerJson(c.Writer, result, err)
}
