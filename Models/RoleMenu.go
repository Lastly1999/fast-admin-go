package Models

// RoleMenu 权限菜单表
type RoleMenu struct {
	MenuId         uint   `json:"menuId" gorm:"column:menuId"`
	MenuName       string `json:"menuName" gorm:"column:menuName"`
	MenuIcon       string `json:"menuIcon" gorm:"column:menuIcon"`
	MenuPath       string `json:"menuPath" gorm:"column:menuPath"`
	MenuParentId   uint   `json:"menuParentId" gorm:"column:menuParentId"`
	MenuParentName string `json:"menuParentName" gorm:"column:menuParentName"`
}
