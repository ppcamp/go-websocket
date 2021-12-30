package app

import (
	"log"
	"net/http"
	"src/internal/config"
	"src/internal/controllers"
	"src/pkg/services/websocket"

	"github.com/urfave/cli/v2"
)

func Run(_ *cli.Context) error {
	config.SetupLoggers()

	ws := websocket.NewServer()
	go ws.Start()
	wrap := func(w http.ResponseWriter, r *http.Request) { controllers.Websocket(ws, w, r) }

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/ws", wrap)

	err := http.ListenAndServe(config.App.Address, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
