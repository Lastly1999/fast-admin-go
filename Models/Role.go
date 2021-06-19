package Models

import "time"

// Role 角色权限表
type Role struct {
	RoleId     uint      `json:"RoleId" gorm:"column:roleId"`
	RoleName   string    `json:"roleName" gorm:"column:roleName"`
	CreateDate time.Time `json:"createDate" gorm:"column:createDate"`
	Describe   string    `json:"describe" gorm:"column:describe"`
}
