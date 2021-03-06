package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// GinConf 配置文件结构体
type GinConf struct {
	Project struct {
		Port string `json:"port"`
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

func ReadYamlConfig() GinConf {
	var ginConf GinConf
	// 	读取项目所在系统路径
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)
	//	读取yaml文件到缓存中
	config, err := ioutil.ReadFile(rootPath + "\\config.yaml")
	if err != nil {
		fmt.Print(err)
	}
	//	yaml文件内容映射到结构体中
	yaml.Unmarshal(config, &ginConf)
	return ginConf
}

func GetConfigFileFromExecutable(fileName string) *os.File {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil
	}
	f, err := os.Open(path.Join(dir, fileName))
	if err != nil {
		return nil
	}
	return f
}
