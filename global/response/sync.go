package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 200
	ERROR   = -200
)

func RequestParamSync(ctx *gin.Context) map[string]interface{} {
	param := make(map[string]interface{})
	err := ctx.BindJSON(&param)
	if err != nil {
		return nil
	}
	return param
}

func JsonResultOk(status bool, data interface{}, ctx *gin.Context) {
	ctx.JSON(SUCCESS, gin.H{
		"status": status,
		"data":   data,
		"msg":    "error",
	})
}

func JsonResultErr(status bool, data interface{}, ctx *gin.Context) {
	ctx.JSON(ERROR, gin.H{
		"status": status,
		"data":   data,
		"msg":    "error",
	})
}
