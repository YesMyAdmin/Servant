package backup

import (
	"butler/internal/model/dto/backup"
	"common/public/pkg"
	backupsvc "butler/internal/service/backup"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 新建备份任务
func NewTask(c *gin.Context) error {
	var req dto.NewBackupTaskReq
	err := c.ShouldBindBodyWithJSON(&req)
	if (err!=nil) {
		return pkg.BadArgumentsError(err.Error())
	}
	err = backupsvc.NewBackupTask(&req)
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
	resp, err := backupsvc.ListTasks(&req);
	if (err!= nil) {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

// 切换任务状态
func SwitchTasks(c *gin.Context) error {
	taskIdStr := c.Request.PathValue("taskId")
	var req dto.SwitchTaskReq
	err := c.ShouldBindJSON(&req)
	
	if (err!=nil) {
		return pkg.BadArgumentsError(err.Error())
	}

	taskId, parseErr := strconv.ParseUint(taskIdStr, 10, 64)
	if (parseErr != nil) {
		return pkg.BadArgumentsError(parseErr.Error())
	}

	backupsvc.BackupTaskSwitch(uint64(taskId), req.Enabled)
	return nil
}

// 删除备份任务
func DeleteBackupTask(c *gin.Context) error {
	taskIdStr := c.Request.PathValue("taskId")
	taskId, parseErr := strconv.ParseUint(taskIdStr, 10, 64)
	if (parseErr != nil) {
		return pkg.BadArgumentsError(parseErr.Error())
	}
	backupsvc.DeleteBackupTask(taskId)
	return nil
}