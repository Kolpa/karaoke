package main

import "net/http"
import "karaoke/internal/server"
import "karaoke/internal/player"


func main() {
    player.Initialize();
	http.HandleFunc("/", server.ServeWs)
	http.ListenAndServe("0.0.0.0:8080", nil)
}