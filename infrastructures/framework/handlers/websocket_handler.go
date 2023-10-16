package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"touchedFlowed/features/user/responses"
	"touchedFlowed/infrastructures/framework/websocket"
)

func UpgradeToWS(c *gin.Context) {
	Ws := websocket.GetInstance()
	socket, err := Ws.HandleWS(c.Writer, c.Request)
	if err != nil {
		return
	}

	user := c.MustGet("user").(*responses.CreateUserResponse)
	socket.SetData("user", user)
	socket.Emit("message", "Hello from server")

	socket.On("join", func(payload string, _ *websocket.Socket) {
		fmt.Printf("%s joined the room: %s\n", user.FirstName, payload)
		socket.Join(payload)
		socket.Broadcast(payload, "message", user.FirstName+" joined the room: "+payload)
	})

	socket.On("leave", func(payload string, _ *websocket.Socket) {
		fmt.Printf("%s left the room: %s\n", user.FirstName, payload)
		socket.Leave(payload)
		socket.Broadcast(payload, "message", user.FirstName+" left the room: "+payload)
	})
	Ws.Listen()
}
