package pkg

import (
	"net/http"
	"time"
	"log/slog"
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
			Error: "common.bad_arguments",
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
	slog.Error(msg)
	return &ServiceError{
		response: ErrorResp{
			Error: "common.internal_error",
			RaiseTime: time.Now(),
			HttpResp: HttpResp{
				Status: http.StatusInternalServerError,
				Msg: msg,
			},
		},
	}
}

// 数据库错误(500)
// tableName 数据表名
// err gorm数据库错误
func DatabaseError(tableName string, err error) error {
	returnMsg := "A server-side error occurs while accessing database. Checking server logs may find the error details."
	slog.Error("Database error occurs", slog.String("tableName", tableName), slog.Any("error", err))
	return &ServiceError{
		response: ErrorResp{
			Error: "common.database_error",
			RaiseTime: time.Now(),
			HttpResp: HttpResp{
				Status: http.StatusInternalServerError,
				Msg: returnMsg,
			},
		},
	}
}

//--------------------业务类型错误--------------------

// 合并文件冲突错误(409)
func FileMergingConflictError(files *[]uint64) error {
	returnMsg := "Files you merging have different types and cannot be merged."
	slog.Error(returnMsg, slog.Any("files", files))
	return &ServiceError{
		response: ErrorResp{
			Error: "butler.backup_files.merge_conflict",
			RaiseTime: time.Now(),
			HttpResp: HttpResp{
				Status: http.StatusConflict,
				Msg: returnMsg,
			},
		},
	}
}