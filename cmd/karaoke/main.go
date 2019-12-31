package main

import "net/http"
import "karaoke/internal/server"
import "karaoke/internal/player"

func main() {
	player.Initialize()

	fs := http.FileServer(http.Dir("page"))

	http.Handle("/", fs)
	http.HandleFunc("/ws", server.ServeWs)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
