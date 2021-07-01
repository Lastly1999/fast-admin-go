package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func InitProjectConfig() {
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
