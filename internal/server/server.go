package server

import "log"
import "net/http"
import "github.com/gorilla/websocket"

import "karaoke/internal/commands"

type remoteCommand struct {
	Command   string
	Arguments interface{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func processCommand(ws *websocket.Conn) {
	var command remoteCommand
	err := ws.ReadJSON(&command)

	if err != nil {
		log.Println("failed to parse message:", err)
		return
	}

	var response interface{}

	switch command.Command {
	case "list_songs":
		response = commands.ListSongs()
	case "add_song":
		response = commands.AddSong(command.Arguments)
	case "get_queue":
		response = commands.GetQueue()
	}

	ws.WriteJSON(response)
}

//ServeWs will respond to websocket requests
func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	for {
		processCommand(ws)
	}
}
