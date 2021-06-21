package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//	全局变量 用来存储剩余的考试时间
var initTimer int64 = 300000

func WsPing(c *gin.Context) {
	// 升级http协议为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	ws.WriteJSON(initTimer)
	if err != nil {
		return
	}
	defer ws.Close()
	//	开启事件循环 对发送过来的数据进行监听写入
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//	每次请求都将时间减掉 但是服务中的全局变量不会发生改变
		initTimer--
		log.Println(initTimer)
		//写入ws数据
		err = ws.WriteMessage(mt, []byte(strconv.FormatInt(initTimer, 10)))
		if err != nil {
			break
		}
	}
}
