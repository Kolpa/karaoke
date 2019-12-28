package server

import "log"
import "net/http"
import "github.com/gorilla/websocket"

import "karaoke/internal/commands"

type RemoteCommand struct {
	Command   string
	Arguments interface{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	for {
		var remoteCommand RemoteCommand
		err := ws.ReadJSON(&remoteCommand)

		if err != nil {
			log.Println("failed to parse message:", err)
			continue
		}

		var response interface{};

		switch remoteCommand.Command {
		case "list_songs":
			response = commands.ListSongs()
		}

		ws.WriteJSON(response)
	}
}
