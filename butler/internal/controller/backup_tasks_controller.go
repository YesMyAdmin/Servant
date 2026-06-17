package controller

import (
	"butler/internal/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询备份任务列表
func listTasks(c *gin.Context) {
	var req dto.ListTasksReq
	req.TaskName = c.Param("taskName");
	pageNum, err := strconv.Atoi(c.Request.PathValue("pageNum"));
	if (err != nil) {
		
	}
	req.PageNum = pageNum

}
