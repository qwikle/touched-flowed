package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"touchedFlowed/infrastructures/framework/websocket"
)

func UpgradeToWS(c *gin.Context) {
	Ws := websocket.GetInstance()
	err := Ws.HandleWS(c.Writer, c.Request)
	if err != nil {
		return
	}
	Ws.On("message", func(data string, socket *websocket.Socket) {
		fmt.Println(data)
		socket.Emit("message", data)
	})
	Ws.OnConnect(func(s *websocket.Socket) {
		fmt.Printf("%s connected\n", s.Id)
		s.Emit("connected", "You are connected")
	})
	Ws.OnDisconnect(func(s *websocket.Socket) {
		fmt.Printf("%s disconnected\n", s.Id)
	})
	Ws.Start()
}
