package main

type Hub struct {
	userMap    *UserMap
	register   chan *User
	unregister chan *User
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		userMap:    NewUserMap(),
		register:   make(chan *User),
		unregister: make(chan *User),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) run() {
	for {
		select {
		case user := <-h.register:
			h.userMap.set(user)
		case user := <-h.unregister:
			h.userMap.del(user.id)
		case msg := <-h.broadcast:
			users := h.userMap.getAll()
			for _, user := range users {
				user.send <- msg
			}
		}
	}
}
