package Models

// User 用户表 gorm:"column:xxx" 为数据库字段映射
type User struct {
	ID         uint   `json:"id,omitempty"`
	UserName   string `json:"userName" gorm:"column:userName"`
	UserPass   string `json:"userPass" gorm:"column:userPass"`
	UserAvatar string `json:"userAvatar" gorm:"column:userAvatar"`
	UserSign   string `json:"userSign" gorm:"column:userSign"`
}
