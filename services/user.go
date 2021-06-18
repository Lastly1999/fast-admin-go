package services

import (
	"fmt"
	"go-service/Models"
	"go-service/global"
)

// FindUserByAuth 查询是否存在用户 登录操作
func FindUserByAuth(userName interface{}, passWord interface{}) (error, Models.User) {
	var user Models.User
	err := global.Db.Where("userName = ? AND userPass = ?", userName, passWord).Find(&user).Error
	fmt.Println(err)
	return err, user
}
