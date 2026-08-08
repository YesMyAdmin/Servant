package butler

import (
	"butler/internal/controller/backup"
	"butler/internal/controller/maids"
	"butler/internal/database"
	"butler/internal/middleware"
	"common/public/config"
	"log"

	"github.com/gin-gonic/gin"
)

//管家节点入口函数
func main(masterPassword string, encryptedConfigPath string) {
	// 解密配置文件
	if (encryptedConfigPath == "") {
		defaultEncryptedConfigPath := "config/config.bin"
		config.DecryptConfig(masterPassword, defaultEncryptedConfigPath)
	} else {
		config.DecryptConfig(masterPassword, encryptedConfigPath)
	}
	
	// 初始化数据库
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	// 创建路由
	r := gin.Default()

	// 注册路由
	// 备份任务管理
	r.GET("/butler/backup/tasks/list", middleware.HandlerFunc(backup.ListTasks))
	r.POST("/butler/backup/tasks/new", middleware.HandlerFunc(backup.NewTask))
	r.POST("/butler/backup/tasks/{taskId}/switch", middleware.HandlerFunc(backup.SwitchTasks))
	r.POST("/butler/backup/tasks/{taskId}/delete", middleware.HandlerFunc(backup.DeleteBackupTask))
	// 备份文件管理
	

	// 女仆节点管理
	r.POST("/butler/maids/register", middleware.HandlerFunc(maids.RegisterMaid))
	r.POST("/butler/maids/{maidId}/dismiss", middleware.HandlerFunc(maids.DismissMaid))
	r.POST("/butler/maids/{maidId}/update", middleware.HandlerFunc(maids.UpdateMaid))
	r.GET("/butler/maids/list", middleware.HandlerFunc(maids.ListMaids))

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
