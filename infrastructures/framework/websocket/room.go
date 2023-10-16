package websocket

type Room struct {
	Name    string
	Sockets map[string]*Socket
}

func (r *Room) AddSocket(s *Socket) {
	r.Sockets[s.Id] = s
}

func (r *Room) RemoveSocket(s *Socket) {
	delete(r.Sockets, s.Id)
}

func NewRoom(name string) *Room {
	return &Room{
		Name:    name,
		Sockets: make(map[string]*Socket),
	}
}
