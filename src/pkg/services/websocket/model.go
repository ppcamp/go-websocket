package websocket

type SocketMessage struct {
	Type     SocketEventType `json:"type"`
	Id       string          `json:"id"`
	Nickname string          `json:"nickname"`
	Message  *string         `json:"message"`
}
