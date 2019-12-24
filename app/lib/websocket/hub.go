package websocket

// PersonalMessage ...
type PersonalMessage struct {
	Client  *Client
	Message []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	Send chan *PersonalMessage

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// NewHub returns new websocket hub.
func NewHub() *Hub {
	return &Hub{
		Send:       make(chan *PersonalMessage),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

// Run starts the hub.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			// logger.Printf("client connected : memberID=[%v]", client.memberID)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
			// logger.Printf("client disconnect : memberID=[%v]", client.memberID)
		case pMessage := <-h.Send:
			c := pMessage.Client
			c.send <- pMessage.Message
			// logger.Printf("personal message : memberID=[%v], msg=[%v]", pMessage.client.memberID, string(pMessage.message))
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.send <- message:
					// log.Printf("client[%v] send message : [%v]\n", client.memberID, string(message))
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
			// logger.Printf("broadcast : msg=[%v]", string(message))
		}
	}
}
