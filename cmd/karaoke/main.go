package main

import "net/http"
import "karaoke/internal/server"
import "karaoke/internal/player"


func main() {
    player.Initialize();
    player.AddSongToQueue("/mnt/c/Users/kolya/Downloads/Rigmar Chart Busters Karaoke CDGMP3/3 Dog Night - Mama Told Me.mp3");
	http.HandleFunc("/", server.ServeWs)
	http.ListenAndServe("0.0.0.0:8080", nil)
}