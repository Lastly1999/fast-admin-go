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
		response.JsonResultOk(true, user, ctx)
	} else {
		response.JsonResultErr(false, nil, ctx)
	}
}
