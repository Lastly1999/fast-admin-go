package router

import (
	"github.com/gin-gonic/gin"
	"go-service/database"
)

// InitRoleRouter 权限路由集合
func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role")
	{
		RoleRouter.GET("/menus/:id", database.RoleHandel.GetMenusByUserIdHandel)
	}
}
