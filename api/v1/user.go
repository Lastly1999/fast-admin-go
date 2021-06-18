package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/global/response"
	"go-service/services"
)

func LoginAction(ctx *gin.Context) {
	param := response.RequestParamSync(ctx)
	userName := param["userName"]
	passWord := param["passWord"]
	err, user := services.FindUserByAuth(userName, passWord)
	fmt.Println(err)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status": 1,
			"msg":    "success!",
			"data":   user,
		})
	} else {
		ctx.JSON(200, gin.H{
			"status": 0,
			"msg":    "error!",
		})
	}
}
