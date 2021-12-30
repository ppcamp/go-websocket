package controllers

import (
	"log"
	"net/http"
	"src/pkg/services/websocket"

	ws "github.com/gorilla/websocket"
)

// change the http server to a websocket connection
var httpToWs = ws.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func Websocket(hub *websocket.Server, w http.ResponseWriter, r *http.Request) {
	conn, err := httpToWs.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := websocket.NewClient(conn, hub)
	hub.RegisterClient(client)
	client.Start()
}
