package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/services"
)

// FindtUserRoleMenus 查询用户权限菜单
func FindtUserRoleMenus(ctx *gin.Context) {
	userId := ctx.Param("id")
	//	查询用户权限菜单
	result := services.FindUserRoleMenus(userId)
	if len(result) > 0 {
		response.JsonResultOk(true, result, ctx)
		return
	}
	response.JsonResultErr(false, ctx)
}
