package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/models"
	"go-service/services"
	"strconv"
)

type RoleHandel struct {
	RoleSrv services.IRoleService
}

// GetMenusByUserIdHandel 查询用户权限菜单
func (r *RoleHandel) GetMenusByUserIdHandel(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		response.JsonResultErr("请求参数有误，请重试", ctx)
		return
	}
	//	强制转换id为int
	id, _ := strconv.Atoi(userId)
	user := models.User{
		ID: id,
	}
	//	查询用户权限菜单
	result, err := r.RoleSrv.ListByUserId(user)
	if err != nil {
		response.JsonResultErr("查询失败，未知错误 role_handel.go-18行", ctx)
	}
	response.JsonResultOk(result, ctx)
}
