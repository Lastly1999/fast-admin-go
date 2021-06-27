package models

import "time"

// User 用户表
type User struct {
	ID         int       `json:"id,omitempty"`
	UserName   string    `json:"userName" gorm:"column:userName"`
	UserPass   string    `json:"userPass" gorm:"column:userPass"`
	UserAvatar string    `json:"userAvatar" gorm:"column:userAvatar"`
	UserSign   string    `json:"userSign" gorm:"column:userSign"`
	CreateDate time.Time `json:"createDate" gorm:"column:createdDate"`
}
