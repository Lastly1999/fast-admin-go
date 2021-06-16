package database

import (
	"fmt"
	"go-service/Models"
	"go-service/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)



func InitDataBase(){
	dsn := "root:1234@tcp(159.75.22.114:3306)/system_manage?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "Admin_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil{
		fmt.Println("database content error...")
	}
	errs := db.AutoMigrate(&Models.User{})
	if errs != nil {
		return 
	}
	global.Db = db
}