package format

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiResultJson struct {
	context *gin.Context
}

func NewApiResult(ctx *gin.Context) *apiResultJson {
	return &apiResultJson{context: ctx}
}

// 设置响应头
func (r *apiResultJson) SetHeader(key, value string) *apiResultJson {
	r.context.Writer.Header().Set(key, value)
	return r
}

// 成功响应
func (r *apiResultJson) Success(data interface{}) {
	apiResult := ApiReulst{
		Success: true,
		Msg:     "成功",
		Obj:     data,
		Code:    "0000",
	}
	r.context.JSON(http.StatusOK, apiResult)
}

// 失败响应
func (r *apiResultJson) Error(errCode string, errMsg string) {
	apiResult := ApiReulst{
		Success: false,
		Msg:     errMsg,
		Obj:     nil,
		Code:    errCode,
	}
	r.context.JSON(http.StatusOK, apiResult)
}

type ApiReulst struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Obj     interface{} `json:"obj"`
	Code    string      `json:"code"`
}
