package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/Global/Response"
	"go-service/Services"
)

// FindtUserRoleMenus 查询用户权限菜单
func FindtUserRoleMenus(ctx *gin.Context) {
	userId := ctx.Query("id")
	result := Services.FindUserRoleMenus(userId)
	Response.JsonResultOk(true, result, ctx)
}
