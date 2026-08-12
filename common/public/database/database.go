package database

import (
	"common/public/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库实例
var DB *gorm.DB

// Connect 从配置里获取连接参数,连接数据库
func Connect() error {
	raw, err := config.Get("general.databases")
	if err != nil {
		return err
	}
	connections, ok := raw.([]interface{})
	// 如果没有配置数据库连接,返回错误
	if !ok || len(connections) == 0 {
		return fmt.Errorf("no database connections configured")
	}
	//获取第一个数据库连接配置,虽然配置文件可以配置多个数据库连接,但是系统目前只支持一个数据库连接
	connection := connections[0]
	// 数据库主机
	host, ok := connection.(map[string]interface{})["host"].(string)
	if !ok {
		return fmt.Errorf("database host configuration 'host' is not configured or invalid")
	}
	// 数据库端口
	port, ok := connection.(map[string]interface{})["port"].(int)
	if !ok {
		return fmt.Errorf("database port configuration 'port' is not configured or invalid")
	}
	// 数据库名称
	database, ok := connection.(map[string]interface{})["database"].(string)
	if !ok {
		return fmt.Errorf("database name configuration 'database' is not configured or invalid")
	}
	// 数据库连接参数
	properties, ok := connection.(map[string]interface{})["properties"].(string)
	if !ok {
		return fmt.Errorf("database properties configuration 'properties' is not configured or invalid")
	}
	//数据库用户
	user, ok := connection.(map[string]interface{})["user"].(string)
	if !ok {
		return fmt.Errorf("database user configuration 'user' is not configured or invalid")
	}
	//用户名不允许使用root和postgres,因为这两个用户权限过高,容易造成安全问题
	if (user == "root" || user == "postgres") {
		return fmt.Errorf("database user configuration 'user' is not allowed to use root or postgres")
	}
	//数据库密码
	password, ok := connection.(map[string]interface{})["password"].(string)
	if !ok {
		return fmt.Errorf("database password configuration 'password' is not configured or invalid")
	}

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?%s`, user, password, host, port, database, properties)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
