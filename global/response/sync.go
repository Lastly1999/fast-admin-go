package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 200
	ERROR   = -200
)

func RequestParamSync(ctx *gin.Context) (map[string]interface{}, error) {
	param := make(map[string]interface{})
	bindErr := ctx.BindJSON(&param)
	if bindErr != nil {
		return nil, errors.New("json数据有误，请重试")
	}
	return param, bindErr
}

// JsonResultOk 成功统一返回格式
func JsonResultOk(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   SUCCESS,
		"status": true,
		"msg":    "success",
		"data":   data,
	})
}

// JsonResultErr 失败统一返回格式
func JsonResultErr(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   ERROR,
		"status": false,
		"msg":    msg,
	})
}

// JsonUploadSuccess 上传成功统一返回格式
func JsonUploadSuccess(data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": true,
		"msg":    msg,
		"data":   data,
	})
}

// JsonUploadErr 上传错误统一返回格式
func JsonUploadErr(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   ERROR,
		"status": true,
		"data":   data,
	})
}
