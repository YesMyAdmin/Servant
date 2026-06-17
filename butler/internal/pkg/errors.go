package pkg

import (
	"net/http"
	"time"
)

type ServiceError struct {
	response ErrorResp
}

func (e ServiceError) Error() string {
	return e.response.Error
}

// HttpStatus 返回 HTTP 状态码
func (e ServiceError) HttpStatus() int {
	return e.response.Status
}

// ErrorResp 返回错误响应体
func (e ServiceError) ErrorResp() ErrorResp {
	return e.response
}

//无效参数错误(400)
func BadArgumentsError(msg string) error {
	return &ServiceError{
		response: ErrorResp{
			Error: "system.bad_arguments",
			RaiseTime: time.Now(),
			HttpResp: HttpResp{
				Status: http.StatusBadRequest,
				Msg: msg,
			},
		},
	}
}

//内部服务器错误(500)
func InternalServerError (msg string) error {
		return &ServiceError{
		response: ErrorResp{
			Error: "system.internal_error",
			RaiseTime: time.Now(),
			HttpResp: HttpResp{
				Status: http.StatusInternalServerError,
				Msg: msg,
			},
		},
	}
}