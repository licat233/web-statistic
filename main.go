package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	// http升级websocket协议的配置
	wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//WebDataMap 用于存储各个网站的在线用户量
	WebDataMap = make(map[string]int)
	//CloseClearMap 通知协程关闭
	CloseClearMap chan byte
)

func main() {
	go ClearMap()
	defer func() {
		close(CloseClearMap)
	}()
	Router := gin.Default()
	Router.GET("/ws/onlineServer", func(c *gin.Context) {
		if c.Query("weburl") == "" {
			c.JSON(200, gin.H{
				"msg": "滚！！！",
			})
			return
		}
		WebSocketHandler(c.Query("weburl"), c.Writer, c.Request)
	})
	if err := Router.Run(":8866"); err != nil {
		fmt.Println("Gin run failed")
		return
	}
}

//WebSocketHandler /ws/onlineServer接口函数
func WebSocketHandler(clientURL string, resp http.ResponseWriter, req *http.Request) {
	WebDataMap[clientURL]++
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	defer func() {
		_ = wsSocket.Close()
		if err := recover(); err != nil {
		}
	}()
	//发送在线人数给客户端，相当于2秒检测一遍用户是否还在
	for {
		if err = wsSocket.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(WebDataMap[clientURL]))); err != nil {
			WebDataMap[clientURL]--
			return
		}
		time.Sleep(time.Second * 2)
	}
}

//ClearMap 协程，用于处理无用的键值对
func ClearMap() {
	for {
		select {
		case <-CloseClearMap:
			fmt.Println("协程已关闭")
			return
		default:
			for k, v := range WebDataMap {
				if v <= 0 {
					delete(WebDataMap, k)
				}
			}
		}
		time.Sleep(time.Minute * 30)
	}
}
