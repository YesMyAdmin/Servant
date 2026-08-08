package router

import (
	"butler/internal/controller/backup"
	"butler/internal/controller/maids"
	"butler/internal/middleware"
	"common/public/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// 启动路由
func Run() error {
	raw, err := config.Get("general.server")
	if err != nil {
		return err
	}
	serverConfig := raw.(interface{})
	//加载端口配置
	port, ok := serverConfig.(map[string]interface{})["port"].(int)
	if !ok {
		return fmt.Errorf("database host configuration 'host' is not configured or invalid")
	}
	r = gin.Default()
	registerBackupTasks()
	registerBackupFiles()
	registerMaids()
	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("Failed to start server: %v", err)
	}
	return nil
}

// 注册备份任务路径
func registerBackupTasks() {
	r.GET("/butler/backup/tasks/list", middleware.HandlerFunc(backup.ListTasks))
	r.POST("/butler/backup/tasks/new", middleware.HandlerFunc(backup.NewTask))
	r.POST("/butler/backup/tasks/{taskId}/switch", middleware.HandlerFunc(backup.SwitchTasks))
	r.POST("/butler/backup/tasks/{taskId}/delete", middleware.HandlerFunc(backup.DeleteBackupTask))
}

// 注册备份文件路径
func registerBackupFiles() {
	r.GET("/butler/backup/files/list", middleware.HandlerFunc(backup.ListBackupFiles))
	r.GET("/butler/backup/files/{fileId}/records", middleware.HandlerFunc(backup.ListBackupFileRecords))
	r.POST("/butler/backup/files/merge", middleware.HandlerFunc(backup.MergeBackupFileRecords))
	r.POST("/butler/backup/files/{backupRecordId}/delete", middleware.HandlerFunc(backup.DeleteBackupRecord))
}

// 注册女仆节点管理路径
func registerMaids() {
	r.POST("/butler/maids/register", middleware.HandlerFunc(maids.RegisterMaid))
	r.POST("/butler/maids/{maidId}/dismiss", middleware.HandlerFunc(maids.DismissMaid))
	r.POST("/butler/maids/{maidId}/update", middleware.HandlerFunc(maids.UpdateMaid))
	r.GET("/butler/maids/list", middleware.HandlerFunc(maids.ListMaids))
}
