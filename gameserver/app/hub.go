package gameserver

// PersonalMessage ...
type PersonalMessage struct {
	client  *Client
	message []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	send chan *PersonalMessage

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		send:       make(chan *PersonalMessage),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			logger.Printf("client connected : memberID=[%v]", client.memberID)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			logger.Printf("client disconnect : memberID=[%v]", client.memberID)
		case pMessage := <-h.send:
			c := pMessage.client
			c.send <- pMessage.message
			logger.Printf("personal message : memberID=[%v], msg=[%v]", pMessage.client.memberID, string(pMessage.message))
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
					// log.Printf("client[%v] send message : [%v]\n", client.memberID, string(message))
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			logger.Printf("broadcast : msg=[%v]", string(message))
		}
	}
}
