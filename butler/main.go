package butler

import (
	"butler/internal/controller"
	"butler/internal/database"
	"butler/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

//管家节点入口函数
func main() {
	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 创建路由
	r := gin.Default()

	// 注册路由
	r.GET("/butler/backup/tasks/list", middleware.HandlerFunc(controller.ListTasks))

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
