package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/services"
)

// GetCarListById 获取车型列表
func GetCarListById(ctx *gin.Context) {
	//	接收参数
	param := response.RequestParamSync(ctx)
	//	类型id
	typeId := param["typeId"].(string)
	//	servcies处理
	result, total := services.FindCarListById(typeId)
	fmt.Println(result)
	fmt.Println(total)
}
