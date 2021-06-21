package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-service/router"
	"time"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	/*
	   cors.New方法返回一个函数参数是c *gin.Context
	   将这个参数赋值给mwCORS,直接当中间间使用,
	   默认修改返回的请求头,实现跨域功能
	   cors.Config为一个结构体,结构体实例后传入cors.New实现生成中间件功能
	*/
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})
	Router.Use(mwCORS)
	apiGroup := Router.Group("v1")
	{
		router.InitUserRouter(apiGroup)
		router.InitRoleRouter(apiGroup)
		router.InitWsRouter(apiGroup)
	}
	return Router
}
