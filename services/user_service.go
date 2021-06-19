package services

import (
	"go-service/global"
	"go-service/models"
)

// FindUserByAuth 查询是否存在用户 登录操作
func FindUserByAuth(userName interface{}, passWord interface{}) (total int64, userInfo models.User) {
	result := global.Db.Where("userName = ? AND userPass = ?", userName, passWord).First(&userInfo)
	total = result.RowsAffected
	return total, userInfo
}
