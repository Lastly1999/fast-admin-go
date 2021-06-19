package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/models"
	"go-service/services"
)

type UserInfoRes struct {
	UserInfo models.User
	RoleId   int `json:"roleId"`
}

// LoginAction 用户登录
func LoginAction(ctx *gin.Context) {
	param := response.RequestParamSync(ctx)
	userName := param["userName"]
	passWord := param["userPass"]
	//	查询用户是否存在
	total, user := services.FindUserByAuth(userName, passWord)
	//	查询用户权限信息
	roleResult := services.FindUserRoleId(user.ID)
	//	返回判断
	if total == 0 {
		response.JsonResultErr(false, ctx)
		return
	}
	response.JsonResultOk(true, UserInfoRes{
		UserInfo: user,
		RoleId:   int(roleResult.RoleId),
	}, ctx)
}
