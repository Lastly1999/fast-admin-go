package Models

// UserRole 用户权限表
type UserRole struct {
	UserId uint `json:"userId" gorm:"column:userId"`
	RoleId uint `json:"roleId" gorm:"column:roleId"`
}
