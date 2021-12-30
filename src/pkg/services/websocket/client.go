package websocket

import (
	"bytes"
	"fmt"
	"src/pkg/helpers"
	"src/pkg/utils"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var logClient = helpers.NewModuleLogger("WebSocketClient")

const (
	writeWait      = 10 * time.Second    // Time allowed to Write a message
	pongWait       = 60 * time.Second    // Time allowed to Read the next pong
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period
	maxMessageSize = bytes.MinRead       // Maximum message size (in bytes)
)

type Client struct {
	socket *websocket.Conn

	Uuid    string
	Nick    string
	addr    string
	Message chan string
	hub     *Server
}

func NewClient(socket *websocket.Conn, server *Server) *Client {
	addr := socket.RemoteAddr()
	uuid := utils.Must(uuid.NewRandom()).(uuid.UUID)
	nick := fmt.Sprintf("AnonymousUser%d", server.UsersLength())

	return &Client{
		Uuid:    strings.Replace(uuid.String(), "-", "", -1),
		socket:  socket,
		hub:     server,
		Message: make(chan string),
		Nick:    nick,
		addr:    utils.Must(helpers.WebsocketAddress(addr)).(string),
	}
}

func (c *Client) changeNick(msg string) {
	contentArray := strings.Split(msg, " ")
	if len(contentArray) >= 2 {
		old := c.Nick
		c.Nick = contentArray[1]
		message := fmt.Sprintf("Client %s changed to %s", old, c.Nick)
		c.hub.Send(NickUpdate, &c.Uuid, &c.Nick, message)
	}
}

// read Opens a thread that will be listening to the
// websocket. It'll be responsible to send the messages got into the Server/hub
func (c *Client) read() {
	defer func() {
		c.hub.UnregisterClient(c)
		c.socket.Close()
	}()

	c.socket.SetReadLimit(maxMessageSize)
	c.socket.SetReadDeadline(time.Now().Add(pongWait))
	c.socket.SetPongHandler(func(string) error { c.socket.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			c.socket.WriteMessage(websocket.TextMessage, []byte("Fail to decode the JSON"))
			return
		}
		msg := string(message)

		// Client/Socket operations
		if utils.StartsWith(msg, OpChangeNick) {
			c.changeNick(msg)
		} else {
			c.hub.Send(Message, &c.Uuid, &c.Nick, msg)
		}
	}
}

// write the messages read from the message channel
// and write it into the socket. This method close the socket connection
// when some errors occurs, therefore, it'll raise an error in the read,
// which will unregister the client
func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.socket.Close()
	}()

	for {
		select {

		case message, ok := <-c.Message:
			if !ok {
				logClient.WithField("client", c.Uuid).Warn("The server closed the channel")
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				logClient.WithField("client", c.Uuid).Warn(err)
				return
			}

		case <-ticker.C:
			c.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				logClient.WithField("client", c.Uuid).
					WithError(err).
					Warn("A ping message was sent and the client didn't respond. Removing the client...")
				return
			}
		}
	}
}

// Start two goroutines to handle with read and write to the websocket clients
// connected
func (c *Client) Start() {
	go c.read()
	go c.write()

	logClient.WithField("client", c.Uuid).Info("Started")
}

// Close the Message channel used to swap messages between clients and server
// management
func (c *Client) Close() {
	logClient.WithField("client", c.Uuid).Info("Closed")
	close(c.Message)
}

// Locals gets the client local port
func (c *Client) Locals() *string { return &c.addr }

// Id gets the client unique id
func (c *Client) Id() *string { return &c.Uuid }
