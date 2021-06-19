package Services

import (
	"go-service/Models"
	"go-service/Global"
)

// FindUserByAuth 查询是否存在用户 登录操作
func FindUserByAuth(userName interface{}, passWord interface{}) (total int64, list interface{}) {
	var userList Models.User
	result := Global.Db.Where("userName = ? AND userPass = ?", userName, passWord).First(&userList)
	total = result.RowsAffected
	return total, userList
}
