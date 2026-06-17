package pkg

import (
	"net/http"
	"time"
)

//分页请求
type PageableReq struct {
	//页号
	PageNum  int `json:"pageNum" form:"pageNum" binding:"required"`
	//页大小
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=5,max=100"`
}

//返回体
type HttpResp struct {
	Status int `json:"status"`
	Msg string `json:"msg,omitempty"`
}

//成功返回体
func SuccessResp(msg string) HttpResp {
	return HttpResp{
		Status: http.StatusOK,
		Msg: msg,
	}
}

//数据返回体
type DataResp[D any] struct {
	Data *D `json:"data,omitempty"`
	HttpResp
}

//错误返回体
type ErrorResp struct {
	//错误编码
	Error string `json:"error"`
	//错误发生的时间
	RaiseTime time.Time `json:"raiseTime"`
	HttpResp
}

//分页返回体
type PageableResp[C any] struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
	Contents []C `json:"contents,omitempty"`
}