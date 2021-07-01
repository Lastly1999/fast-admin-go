package database

import (
	"fmt"
	"github.com/spf13/viper"
	v1 "go-service/handler/v1"
	"go-service/repository"
	"go-service/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB         *gorm.DB
	UserHandel v1.UserHandel
	RoleHandel v1.RoleHandel
)

//	InitProjectConf 初始化项目配置
func InitProjectConf() {
	//	读取yaml数据
	viper.SetConfigName("config")
	//	设置配置文件和可执行文件在同一个目录
	viper.AddConfigPath(".")
	//	异常处理
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
	}
	fmt.Println("read file success..")
}

//	InitDataBase 初始化项目配置
func InitDataBase() {
	fmt.Println()
	//	读取yaml数据库配置信息
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")
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
	DB = db
}

func InitHandel() {
	UserHandel = v1.UserHandel{
		UserSrv: &services.UserService{
			UserRepo: &repository.UserRepository{
				DB: DB,
			},
			RoleRepo: &repository.RolePository{
				DB: DB,
			},
		},
	}
	RoleHandel = v1.RoleHandel{
		RoleSrv: &services.RoleService{
			Repo: &repository.RolePository{
				DB: DB,
			},
		},
	}
}

func init() {
	//	读取配置文件
	InitProjectConf()
	//	初始化数据库
	InitDataBase()
	//	初始化处理器
	InitHandel()
}
