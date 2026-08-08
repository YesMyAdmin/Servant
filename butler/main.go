package butler

import (
	"butler/internal/database"
	"butler/internal/router"
	"common/public/config"
	"log/slog"
	"os"
)

//管家节点入口函数
func main(masterPassword string, encryptedConfigPath string) {
	if("" == masterPassword) {
		slog.Error("Master password of encrypted config file must not be empty.")
		os.Exit(1)
	}
	var path string
	if (encryptedConfigPath == "") {
		path = "config/config.encryped.bin"
	} else {
		path = encryptedConfigPath
	}
	// 解密配置文件
	if configErr := config.DecryptConfig(masterPassword, path); configErr != nil {
		slog.Error(configErr.Error())
		os.Exit(1)
	}
	// 初始化数据库
	if err := database.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	//启动路由
	if err := router.Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
