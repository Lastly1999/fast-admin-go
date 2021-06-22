package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

//	CreateFileUpload 写入文件流主要io.Copy
func CreateFileUpload(file multipart.File, header *multipart.FileHeader) (fileErr error, filePath string) {
	//	获取项目根路径
	rootPath, _ := os.Getwd()
	//	获取请求的文件名称
	fileName := header.Filename
	//	创建文件，如果存在，则会覆盖
	out, _ := os.Create(rootPath + "/file/" + fileName)
	defer out.Close()
	//	创建文件 这边使用了io.Copy来写入文件，是在文件指针之间复制的 因此不会造成文件过大造成内存泄露
	_, fileErr = io.Copy(out, file)
	filePath = "http://localhost:8080/file/" + fileName
	return fileErr, filePath
}

//	ReadServerFile 读取服务器中所有的文件路径及名称
func ReadServerFile() {
	rootPath, _ := os.Getwd()
	file, err := os.ReadFile(rootPath + "/file")
	if err != nil {
		return
	}
	fmt.Println(file)
}
