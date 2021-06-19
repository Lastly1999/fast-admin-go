package Models

type Permission struct {
	ID           uint `json:"id" gorm:"column:id"`
	PermissionId uint `json:"permissionId" gorm:" permissionId"`
	RoleId       uint `json:"roleId" gorm:"column:roleId"`
}
