package DataBase

import (
	"fmt"
	"go-service/Global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDataBase() {
	dsn := "root:Chen1027@tcp(sh-cynosdbmysql-grp-gtd0epo6.sql.tencentcdb.com:28856)/system_manage?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "admin_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("DataBase content error...")
	} else {
		fmt.Println("DataBase content success...")
	}
	Global.Db = db
}
