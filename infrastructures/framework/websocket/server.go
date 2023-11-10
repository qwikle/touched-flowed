package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type SocketServer interface {
	HandleWS(writer http.ResponseWriter, request *http.Request) error
	Emit(roomName string, data string, socket *Socket)
	Broadcast(event string, data string)
	AddSocket(s *Socket)
	RemoveSocket(s *Socket)
	AddSocketToRoom(roomName string, s *Socket)
	CreateRoom(roomName string) *Room
	GetRoom(roomName string) *Room
	RemoveSocketFromRoom(roomName string, s *Socket)
	On(event string, callback SocketEventCallback)
	OnConnect(callback Callback)
	OnDisconnect(callback Callback)
	Start()
	Close()
}

type Callback func(*Socket)

func stringify(event string, data string) (string, error) {
	var socketMessage SocketMessage
	socketMessage.Event = event
	socketMessage.Payload = data
	message, err := json.Marshal(socketMessage)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func parse(message string) (SocketMessage, error) {
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
	rooms             sync.Map
	sockets           sync.Map
	connectedSockets  chan *Socket
	leavedSockets     chan *Socket
	socketChannel     chan *Socket
	events            sync.Map
	mu                sync.Mutex
	onConnect         Callback
	onConnectIsSet    bool
	onDisconnect      Callback
	onDisconnectIsSet bool
}

func (ss *socketServer) On(event string, callback SocketEventCallback) {
	_, ok := ss.events.Load(event)
	if !ok {
		ss.events.Store(event, callback)
	}
}

func (ss *socketServer) Start() {
	for {
		select {
		case socket := <-ss.connectedSockets:
			if ss.onConnect != nil {
				ss.onConnect(socket)
			}
			ss.socketChannel <- socket
		case socket := <-ss.leavedSockets:
			if ss.onDisconnect != nil {
				ss.onDisconnect(socket)
			}
		case socket := <-ss.socketChannel:
			go func() {
				ss.events.Range(func(key, value interface{}) bool {
					event := key.(string)
					callback := value.(SocketEventCallback)
					socket.On(event, callback)
					return true
				})
			}()
			socket.Listen()
		}
	}
}

func (ss *socketServer) HandleWS(writer http.ResponseWriter, request *http.Request) error {
	ws, err := upgrade.Upgrade(writer, request, nil)
	if err != nil {
		return err
	}
	s := NewSocket(ws, request.RemoteAddr)
	ss.AddSocket(s)
	return nil
}

func (ss *socketServer) AddSocket(s *Socket) {
	ss.sockets.Store(s.Id, s)
	ss.connectedSockets <- s
}

func (ss *socketServer) OnConnect(callback Callback) {
	if ss.onConnectIsSet {
		log.Panic("OnConnect callback can be set only once")
		return
	}
	ss.onConnect = callback
}

func (ss *socketServer) OnDisconnect(callback Callback) {
	if ss.onDisconnectIsSet {
		log.Panic("OnDisconnect callback can be set only once")
		return
	}
	ss.onDisconnect = callback
}

func (ss *socketServer) RemoveSocket(s *Socket) {
	ss.rooms.Range(func(key, value interface{}) bool {
		room := value.(*Room)
		room.RemoveSocket(s)
		return true
	})

	ss.sockets.Delete(s.Id)
	ss.leavedSockets <- s
}

func (ss *socketServer) AddSocketToRoom(roomName string, s *Socket) {
	room := ss.GetRoom(roomName)
	if room == nil {
		room = ss.CreateRoom(roomName)
	}
	room.AddSocket(s)
}

func (ss *socketServer) RemoveSocketFromRoom(roomName string, s *Socket) {
	room := ss.GetRoom(roomName)
	if room == nil {
		return
	}
	room.RemoveSocket(s)
}

func (ss *socketServer) Emit(roomName string, data string, s *Socket) {
	socketMessage, err := parse(data)
	if err != nil {
		return
	}
	room := ss.GetRoom(roomName)
	if room == nil {
		log.Panicf("Room %s does not exist", roomName)
		return
	}
	room.Sockets.Range(func(key, value interface{}) bool {
		socket := value.(*Socket)
		if socket.Id != s.Id {
			socket.Emit(socketMessage.Event, data)
		}
		return true
	})
}

func (ss *socketServer) Broadcast(event string, data string) {
	message, err := stringify(event, data)
	if err != nil {
		return
	}
	ss.sockets.Range(func(key, value interface{}) bool {
		socket := value.(*Socket)
		socket.Emit(event, message)
		return true
	})
}

func (ss *socketServer) CreateRoom(roomName string) *Room {
	if room := ss.GetRoom(roomName); room == nil {
		ss.rooms.Store(roomName, NewRoom(roomName))
	}
	room, _ := ss.rooms.Load(roomName)
	return room.(*Room)
}

func (ss *socketServer) GetRoom(roomName string) *Room {
	room, ok := ss.rooms.Load(roomName)
	if !ok {
		return nil
	}
	return room.(*Room)
}

func (ss *socketServer) Close() {
	ss.sockets.Range(func(key, value interface{}) bool {
		socket := value.(*Socket)
		socket.Close()
		return true
	})
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
			connectedSockets: make(chan *Socket, 100),
			leavedSockets:    make(chan *Socket, 100),
			socketChannel:    make(chan *Socket, 100),
		}
	}
	return socketServerInstance
}
