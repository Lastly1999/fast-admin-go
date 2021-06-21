package models

// RoleMenu 权限菜单表
type RoleMenu struct {
	ID             int    `json:"id" gorm:"menuId"`
	MenuId         int    `json:"key" gorm:"column:menuId"`
	MenuName       string `json:"name" gorm:"column:menuName"`
	MenuIcon       string `json:"icon" gorm:"column:menuIcon"`
	MenuPath       string `json:"path" gorm:"column:menuPath"`
	MenuParentId   int    `json:"pid" gorm:"column:menuParentId"`
	MenuParentName string `json:"pName" gorm:"column:menuParentName"`
}

// SyncRoleMenuList 序列化后的集合结构体
type SyncRoleMenuList []RoleMenu

// SyncRoleItem 序列化后的集合单元结构
type SyncRoleItem struct {
	RoleMenu
	Children []SyncRoleItem `json:"children"`
}

// ProcessToTree 序列化数据为单元树形结构
func (m *SyncRoleMenuList) ProcessToTree(pid int, level int) []SyncRoleItem {
	var menuTree []SyncRoleItem
	if level == 10 {
		return menuTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return menuTree
	}

	for _, v := range list {
		child := m.ProcessToTree(v.MenuId, level+1)
		menuTree = append(menuTree, SyncRoleItem{v, child})
	}

	return menuTree
}

//	findChildren 递归遍历子项
func (m *SyncRoleMenuList) findChildren(pid int) []RoleMenu {
	var child []RoleMenu

	for _, v := range *m {
		if v.MenuParentId == pid {
			child = append(child, v)
		}
	}
	return child
}
