package repository

import (
	"go-service/models"
	"gorm.io/gorm"
)

type RolePository struct {
	DB *gorm.DB
}

type IRoleRepository interface {
	ListByUserId(user models.User) (*[]models.SyncRoleItem, error)
}

func (r *RolePository) ListByUserId(user models.User) (*[]models.SyncRoleItem, error) {
	var roleMenus []models.RoleMenu
	sql := "SELECT * FROM admin_user as a LEFT JOIN admin_user_role as b ON a.id = b.userId LEFT JOIN admin_role as c ON b.roleId = c.roleId LEFT JOIN admin_permission as d ON c.roleId = d.roleId LEFT JOIN admin_role_menu as e ON d.permissionId = e.menuId WHERE a.id = ?"
	err := r.DB.Raw(sql, user.ID).Scan(&roleMenus).Error
	// 设置[]SyncRoleMenuList初始数据
	ml := models.SyncRoleMenuList(roleMenus)
	// 生成完全树
	result := ml.ProcessToTree(0, 0)
	return &result, err
}
