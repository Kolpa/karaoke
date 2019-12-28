package main

import "net/http"
import "karaoke/internal/server"


func main() {
	http.HandleFunc("/", server.ServeWs)
	http.ListenAndServe("0.0.0.0:8080", nil)
}