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
		ctx.JSON(200, gin.H{
			"status": false,
			"code":   ERROR,
			"msg":    "获取JSON参数失败，Func -> RequestParamSync",
		})
		return nil
	}
	return param
}

func JsonResultOk(status bool, data interface{}, ctx *gin.Context) {
	ctx.JSON(SUCCESS, gin.H{
		"code":   SUCCESS,
		"status": status,
		"msg":    "success",
		"data":   data,
	})
}

func JsonResultErr(status bool, ctx *gin.Context) {
	ctx.JSON(ERROR, gin.H{
		"code":   ERROR,
		"status": status,
		"msg":    "error",
	})
}
