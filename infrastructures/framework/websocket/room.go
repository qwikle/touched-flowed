package websocket

import (
	"sync"
	"sync/atomic"
)

type Room struct {
	Name    string
	Sockets sync.Map
	count   uint64
}

func (r *Room) AddSocket(s *Socket) {
	r.Sockets.Store(s.Id, s)
	atomic.AddUint64(&r.count, 1)
}

func (r *Room) RemoveSocket(s *Socket) {
	r.Sockets.Delete(s.Id)
	atomic.AddUint64(&r.count, ^uint64(0))
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}

func (r *Room) Length() uint64 {
	return atomic.LoadUint64(&r.count)
}
