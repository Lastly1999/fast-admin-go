package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/Models"
	"go-service/global"
)

// Test 测试
func Test(ctx *gin.Context) {
	var user []Models.User
	global.Db.First(&user)
	fmt.Println(user)
	ctx.JSON(200, gin.H{
		"query":  ctx.Query("id"),
		"result": user,
	})
}

// PostBody 提交
func PostBody(ctx *gin.Context) {
	body := make(map[string]interface{})
	err := ctx.BindJSON(&body)
	if err != nil {
		return
	}
	// 接受json参数
	ctx.JSON(200, gin.H{
		"body": body,
	})
}
