package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/services"
	"strconv"
)

// GetCarListById 获取车型列表
func GetCarListById(ctx *gin.Context) {
	//	接收参数
	param := response.RequestParamSync(ctx)
	//	类型id
	typeId := param["typeId"].(string)
	//	分页参数 页码
	page, _ := strconv.Atoi(param["page"].(string))
	//	分页参数 页数据量
	pageSize, _ := strconv.Atoi(param["pageSize"].(string))
	fmt.Println(pageSize)
	//	servcies处理
	err, result, pageResult := services.FindCarListById(typeId, page, pageSize)

	if err != nil {
		response.JsonResultErr("查询失败", ctx)
	} else {
		response.JsonResultOk(gin.H{
			"list":     result,
			"pageInfo": pageResult,
		}, ctx)
	}
}
