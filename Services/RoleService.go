package Services

import (
	"go-service/Models"
	"go-service/Global"
)

// FindUserRoleMenus 用户id查询用户权限菜单
func FindUserRoleMenus(userId interface{}) (roleMenus []Models.RoleMenu) {
	sql := "SELECT * FROM admin_user as a LEFT JOIN admin_user_role as b ON a.id = b.userId LEFT JOIN admin_role as c ON b.roleId = c.roleId LEFT JOIN admin_permission as d ON c.roleId = d.roleId LEFT JOIN admin_role_menu as e ON d.permissionId = e.menuId WHERE a.id = ?"
	Global.Db.Raw(sql, userId).Scan(&roleMenus)
	return roleMenus
}
