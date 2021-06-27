package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/services"
)

// UploadFile 文件上传
func UploadFile(ctx *gin.Context) {
	//	获取上传请求的头文件
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println("upload file error")
	}
	result, _ := services.CreateFileUpload(file, header)
	response.JsonUploadSuccess(result, "上传成功", ctx)
}
