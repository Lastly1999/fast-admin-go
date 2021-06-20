package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/api/v1"
)

// InitRoleRouter 权限路由集合
func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role")
	{
		RoleRouter.GET("/menus", v1.FindtUserRoleMenus)
	}
}
