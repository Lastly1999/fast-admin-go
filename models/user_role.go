package models

// UserRole 用户权限表
type UserRole struct {
	UserId uint `json:"userId" gorm:"column:userId"`
	RoleId uint `json:"roleId" gorm:"column:roleId"`
}

func (UserRole) TableName() string {
	return "admin_user_role"
}
