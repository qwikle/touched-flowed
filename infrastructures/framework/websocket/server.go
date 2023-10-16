package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
)

type SocketServer interface {
	HandleWS(writer http.ResponseWriter, request *http.Request) (*Socket, error)
	Emit(roomName string, data string, socket *Socket)
	Broadcast(event string, data string)
	AddSocket(s *Socket)
	RemoveSocket(s *Socket)
	AddSocketToRoom(roomName string, s *Socket)
	CreateRoom(roomName string)
	GetRoom(roomName string) *Room
	RemoveSocketFromRoom(roomName string, s *Socket)
	Listen()
	Close()
}

func eventToString(event string, data string) (string, error) {
	var socketMessage SocketMessage
	socketMessage.Event = event
	socketMessage.Payload = data
	message, err := json.Marshal(socketMessage)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func stringToEvent(message string) (SocketMessage, error) {
	var socketMessage SocketMessage
	err := json.Unmarshal([]byte(message), &socketMessage)
	if err != nil {
		return SocketMessage{}, err
	}
	return socketMessage, nil
}

var socketServerInstance SocketServer

var upgrade = websocket.Upgrader{
	ReadBufferSize:  4092,
	WriteBufferSize: 4092,
}

type socketServer struct {
	rooms          map[string]*Room
	sockets        map[string]*Socket
	socketsChannel chan *Socket
}

func (ss *socketServer) Listen() {
	for {
		select {
		case socket := <-ss.socketsChannel:
			socket.Listen()
		}
	}
}

func (ss *socketServer) HandleWS(writer http.ResponseWriter, request *http.Request) (*Socket, error) {
	ws, err := upgrade.Upgrade(writer, request, nil)
	if err != nil {
		return nil, err
	}
	s := NewSocket(ws, request.RemoteAddr)
	ss.AddSocket(s)
	return s, nil
}

func (ss *socketServer) AddSocket(s *Socket) {
	ss.sockets[s.Id] = s
	ss.socketsChannel <- s
}

func (ss *socketServer) RemoveSocket(s *Socket) {
	for _, room := range ss.rooms {
		delete(room.Sockets, s.Id)
	}
	delete(ss.sockets, s.Id)
}

func (ss *socketServer) AddSocketToRoom(roomName string, s *Socket) {
	room := ss.GetRoom(roomName)
	if room == nil {
		ss.CreateRoom(roomName)
		room = ss.GetRoom(roomName)
	}
	room.Sockets[s.Id] = s
}

func (ss *socketServer) RemoveSocketFromRoom(roomName string, s *Socket) {
	if room, o := ss.rooms[roomName]; o {
		delete(room.Sockets, s.Id)
		message, err := eventToString("leave", s.Id)
		if err != nil {
			return
		}
		ss.Emit(roomName, message, nil)
	}
}

func (ss *socketServer) Emit(roomName string, data string, s *Socket) {
	socketMessage, err := stringToEvent(data)
	if err != nil {
		return
	}
	room := ss.GetRoom(roomName)
	if room == nil {
		return
	}
	for _, socket := range room.Sockets {
		if s == nil || socket.Id != s.Id {
			socket.Emit(socketMessage.Event, socketMessage.Payload)
		}
	}
}

func (ss *socketServer) Broadcast(event string, data string) {
	message, err := eventToString(event, data)
	if err != nil {
		return
	}
	for _, socket := range ss.sockets {
		socket.Emit(event, message)
	}
}

func (ss *socketServer) CreateRoom(roomName string) {
	if room := ss.GetRoom(roomName); room == nil {
		ss.rooms[roomName] = NewRoom(roomName)
	}
}

func (ss *socketServer) GetRoom(roomName string) *Room {
	return ss.rooms[roomName]
}

func (ss *socketServer) Close() {
	for _, socket := range ss.sockets {
		socket.Close()
	}
}

func GetInstance() SocketServer {
	if socketServerInstance == nil {
		socketServerInstance = newSocketServer()
	}
	return socketServerInstance
}

func newSocketServer() SocketServer {
	if socketServerInstance == nil {
		socketServerInstance = &socketServer{
			rooms:          make(map[string]*Room),
			sockets:        make(map[string]*Socket),
			socketsChannel: make(chan *Socket, 100),
		}
	}
	return socketServerInstance
}
