package response

import (
	"net/http"

	"github.com/iokr/bbs/core/render"
	"github.com/iokr/bbs/internal/errors"
)

func ServerJson(w http.ResponseWriter, result interface{}, err error) {
	resp := new(Response)
	resp.Result = result
	if err == nil {
		resp.Code = int(errors.Success)
		resp.Message = "success"
		httpRender(w, render.JSON{Data: resp})
		return
	}
	resp.Message = err.Error()
	if value, ok := err.(errors.Error); ok {
		resp.Code = int(value)
	} else {
		resp.Code = int(errors.Default)
	}
	httpRender(w, render.JSON{Data: resp})
}

func httpRender(w http.ResponseWriter, r render.Render) {
	w.WriteHeader(http.StatusOK)
	r.Render(w)
}
