package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// 使用mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateConnection 建立连接
func CreateConnection() (*gorm.DB, error) {

	// 从系统环境变量获取数据库信息
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, DBName,
		),
	)
}
