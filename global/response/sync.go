package response

import "github.com/gin-gonic/gin"

func RequestParamSync(ctx *gin.Context) map[string]interface{} {
	param := make(map[string]interface{})
	err := ctx.BindJSON(&param)
	if err != nil {
		return nil
	}
	return param
}
