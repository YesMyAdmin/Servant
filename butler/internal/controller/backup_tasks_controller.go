package controller

import (
	"butler/internal/model/dto"
	"butler/internal/pkg"
	"butler/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


// 新建备份任务
func NewTask(c *gin.Context) error {
	var req dto.NewBackupTaskReq
	err := c.ShouldBindBodyWithJSON(&req)
	if (err!=nil) {
		return pkg.BadArgumentsError(err.Error())
	}
	err = service.NewBackupTask(&req)
	if (nil != err) {
		return err
	}
	c.JSON(http.StatusOK, pkg.SuccessMessageResp(""))
	return nil
}

//查询备份任务列表
func ListTasks(c *gin.Context) error {
	var req dto.ListTasksReq
	err := c.ShouldBindQuery(&req)
	if (err!=nil) {
		return pkg.BadArgumentsError(err.Error())
	}
	resp, err := service.ListTasks(&req);
	if (err!= nil) {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

