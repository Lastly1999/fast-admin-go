package database

import (
	"fmt"
	"go-service/global"
	"go-service/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDataBase() {
	//	读取yaml数据库配置信息
	dbConfig := utils.ReadYamlConfgi()
	username := dbConfig.Mysql.Username
	password := dbConfig.Mysql.Password
	host := dbConfig.Mysql.Host
	port := dbConfig.Mysql.Port
	database := dbConfig.Mysql.Database
	charset := dbConfig.Mysql.Charset
	//	读取数据库连接参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	//	建立gorm连接池
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "admin_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("database content error...")
		return
	}
	db.Logger.LogMode(0)
	//	赋值全局连接池参数
	global.Db = db
}
