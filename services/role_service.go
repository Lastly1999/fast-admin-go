package services

import (
	"go-service/models"
	"go-service/repository"
)

type IRoleService interface {
	ListByUserId(user models.User) (*[]models.SyncRoleItem, error)
}

type RoleService struct {
	Repo repository.RolePository
}

// ListByUserId 用户id查询用户权限菜单
func (r *RoleService) ListByUserId(user models.User) (*[]models.SyncRoleItem, error) {
	return r.Repo.ListByUserId(user)
}

// FindUserRoleMenus 用户id查询用户权限菜单
//func FindUserRoleMenus(userId interface{}) ([]models.SyncRoleItem, error) {
//	var roleMenus []models.RoleMenu
//	sql := "SELECT * FROM admin_user as a LEFT JOIN admin_user_role as b ON a.id = b.userId LEFT JOIN admin_role as c ON b.roleId = c.roleId LEFT JOIN admin_permission as d ON c.roleId = d.roleId LEFT JOIN admin_role_menu as e ON d.permissionId = e.menuId WHERE a.id = ?"
//	err := global.Db.Raw(sql, userId).Scan(&roleMenus).Error
//	// 设置[]SyncRoleMenuList初始数据
//	ml := models.SyncRoleMenuList(roleMenus)
//	// 生成完全树
//	result := ml.ProcessToTree(0, 0)
//	return result, err
//}

// FindUserRoleId 查询用户的权限id
//func FindUserRoleId(userId interface{}) (userRole models.UserRole) {
//	global.Db.Where("userId = ?", userId).First(&userRole)
//	fmt.Println(userRole)
//	return userRole
//}
