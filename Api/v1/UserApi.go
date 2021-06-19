package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/Global/Response"
	"go-service/Services"
)

// LoginAction 用户登录
func LoginAction(ctx *gin.Context) {
	param := Response.RequestParamSync(ctx)
	userName := param["userName"]
	passWord := param["passWord"]
	total, user := Services.FindUserByAuth(userName, passWord)
	if total == 0 {
		Response.JsonResultErr(false, ctx)
		return
	}
	Response.JsonResultOk(true, user, ctx)
}
