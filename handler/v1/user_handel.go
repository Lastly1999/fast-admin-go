package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/models"
	"go-service/services"
)

type UserHandel struct {
	UserSrv services.IUserService
}

// GetUserHandel 获取用户
func (srv *UserHandel) GetUserHandel(ctx *gin.Context) {
	param, _ := response.RequestParamSync(ctx)
	//	参数传递异常
	if param["userName"] == nil || param["userPass"] == nil {
		response.JsonResultErr("json数据有误，请重试", ctx)
		return
	}
	user := models.User{
		UserName: param["userName"].(string),
		UserPass: param["userPass"].(string),
	}
	result, srvErr := srv.UserSrv.Get(user)
	//	数据库的未知错误捕获
	if srvErr != nil {
		response.JsonResultErr("数据库未知错误，user_handel.go-33行", ctx)
		return
	}
	response.JsonResultOk(gin.H{
		"userInfo": result,
	}, ctx)
}
