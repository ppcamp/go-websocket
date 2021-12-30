package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"src/internal/app"
	"src/internal/config"
)

func main() {
	application := cli.NewApp()
	application.Name = "go-websocket"
	application.Description = "A Golang simple websocket server"
	application.Usage = "go-websocket server || go-websocket client 'name'"
	application.Flags = config.Flags
	application.Action = app.Run
	application.Run(os.Args)
}
