package websocket

import (
	"encoding/json"
	"src/pkg/helpers"
	"src/pkg/utils"

	"github.com/sirupsen/logrus"
)

type Server struct {
	clients map[*string]*Client

	// messages that will be sent to the others (all)
	broadcast chan string

	// register a new client
	register chan *Client

	// unregister a client
	unregister chan *Client

	log *logrus.Entry
}

// NewServer create and returns a new websocket server
func NewServer() *Server {
	return &Server{
		clients:    make(map[*string]*Client),
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		log:        helpers.NewModuleLogger("WebSocketServer"),
	}
}

func (s *Server) addClient(c *Client) {
	s.clients[c.Id()] = c
}

func (s *Server) removeClient(c *Client) {
	if _, ok := s.clients[c.Id()]; ok {
		delete(s.clients, c.Id())
		c.Close()
	}
}

func (s *Server) RegisterClient(c *Client) {
	s.log.Infof("client %s registered", *c.Id())
	s.register <- c
}
func (s *Server) UnregisterClient(c *Client) {
	s.log.Infof("client %s removed", *c.Id())
	message := c.Nick + " has disconnected"
	s.unregister <- c
	s.Send(Notification, &c.Uuid, &c.Nick, message)
}

// Start the websocket server
// TODO: must implement a stop to the server
//
// Example:
//	ws := websocket.NewServer()
//	go ws.Start()
func (s *Server) Start() {
	for {
		select {
		case client := <-s.register:
			s.addClient(client)

		case client := <-s.unregister:
			s.removeClient(client)

		case message := <-s.broadcast:
			// broadcast to clients
			for name := range s.clients {
				select {
				case s.clients[name].Message <- message:
				default:
					s.removeClient(s.clients[name])
				}
			}
		}
	}
}

func (s *Server) Send(t SocketEventType, clientId, nick *string, message string) {
	obj := utils.Must(json.Marshal(SocketMessage{t, *clientId, *nick, &message})).([]byte)
	s.broadcast <- string(obj)
}

func (s *Server) UsersLength() int { return len(s.clients) }
