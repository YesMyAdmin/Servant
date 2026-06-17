package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库实例
var DB *gorm.DB

// Init 初始化数据库连接，从环境变量 MYSQL_DSN 读取连接串
func Init() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:123456@tcp(127.0.0.1:3306)/yesmysql?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
