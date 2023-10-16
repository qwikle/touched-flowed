package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type SocketEventCallback func(data string, socket *Socket)

type SocketHandler struct {
	event    string
	callback SocketEventCallback
}

type Socket struct {
	conn            *websocket.Conn
	Id              string
	Ip              string
	mu              sync.Mutex
	IncomingMessage chan string
	OutgoingMessage chan string
	events          map[string]SocketHandler
	Payload         map[string]interface{}
}

func (s *Socket) SetData(key string, value interface{}) {
	s.Payload[key] = value
}

func (s *Socket) GetData(key string) interface{} {
	return s.Payload[key]
}

func (s *Socket) Emit(event string, data string) {
	message, err := eventToString(event, data)
	if err != nil {
		return
	}
	s.OutgoingMessage <- message
}

func (s *Socket) listenOutgoingMessage() {
	for message := range s.OutgoingMessage {
		err := s.conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println(err)
			s.Close()
			return
		}
	}
}

func (s *Socket) listenIncomingMessage() {
	for {
		select {
		case message, ok := <-s.IncomingMessage:
			if !ok {
				return
			}
			var socketMessage SocketMessage
			err := json.Unmarshal([]byte(message), &socketMessage)
			if err != nil {
				return
			}
			if _, exists := s.events[socketMessage.Event]; exists {
				s.events[socketMessage.Event].callback(socketMessage.Payload, s)
			}
		}
	}
}

func (s *Socket) readMessage() {
	for {
		var socketMessage SocketMessage
		err := s.conn.ReadJSON(&socketMessage)
		if err != nil {
			fmt.Println(err)
			s.Close()
			return
		}
		if _, exists := s.events[socketMessage.Event]; exists {
			message, err := json.Marshal(socketMessage)
			if err != nil {
				fmt.Println(err)
				continue
			}
			s.IncomingMessage <- string(message)
		}
	}
}

func (s *Socket) Listen() {
	go s.listenOutgoingMessage()
	go s.listenIncomingMessage()
	s.readMessage()
}

type SocketMessage struct {
	Event   string `json:"event"`
	Payload string `json:"payload"`
}

func (s *Socket) Join(room string) {
	GetInstance().AddSocketToRoom(room, s)
}

func (s *Socket) Leave(room string) {
	GetInstance().RemoveSocketFromRoom(room, s)
}

func (s *Socket) Broadcast(roomName string, event string, data string) {
	message, err := eventToString(event, data)
	if err != nil {
		return
	}
	GetInstance().Emit(roomName, message, s)
}

func (s *Socket) BroadcastToAll(event string, data string) {
	GetInstance().Broadcast(event, data)
}

func (s *Socket) LeaveAll() {
	GetInstance().RemoveSocket(s)
}

func (s *Socket) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.LeaveAll()
	err := s.conn.Close()
	if err != nil {
		panic(err)
	}
	close(s.IncomingMessage)
	close(s.OutgoingMessage)
}

func (s *Socket) On(event string, callback SocketEventCallback) {
	s.events[event] = SocketHandler{
		event:    event,
		callback: callback,
	}
}

func NewSocket(conn *websocket.Conn, ip string) *Socket {
	return &Socket{
		conn:            conn,
		Ip:              ip,
		Id:              uuid.Must(uuid.NewRandom()).String(),
		IncomingMessage: make(chan string, 100),
		OutgoingMessage: make(chan string, 100),
		events:          make(map[string]SocketHandler),
		Payload:         make(map[string]interface{}),
	}

}
