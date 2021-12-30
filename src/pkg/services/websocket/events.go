package websocket

type SocketEventType string

const (
	Message      SocketEventType = "message"
	Notification SocketEventType = "notification"
	NickUpdate   SocketEventType = "nick_update"
)

const (
	OpChangeNick = "/nick"
)
