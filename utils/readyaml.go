package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// GinConf 配置文件结构体
type GinConf struct {
	Project struct {
		Port int `json:"port"`
	} `json:"project"`
	Mysql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
		Charset  string `json:"charset"`
	} `json:"mysql"`
}

func ReadYamlConfgi() GinConf {
	var ginConf GinConf
	// 	读取项目所在系统路径
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)
	//	读取yaml文件到缓存中
	config, err := ioutil.ReadFile(rootPath + "\\config\\config.yaml")
	if err != nil {
		fmt.Print(err)
	}
	//	yaml文件内容映射到结构体中
	yaml.Unmarshal(config, &ginConf)
	return ginConf
}
