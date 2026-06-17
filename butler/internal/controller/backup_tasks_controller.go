package controller

import (
	"butler/internal/dto"
	"butler/internal/pkg"
	"butler/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
