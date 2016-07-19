package main

import (
	"fmt"
	"mmover/server/config"
	"mmover/server/connection"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	// Redirecting requests to resources' folder
	http.Handle("/", http.FileServer(http.Dir("../client")))

	// data needed from the device app
	http.HandleFunc("/data", connection.GetData)

	// Activate ws handler to retrieve browser data
	http.Handle("/socket", websocket.Handler(connection.WebSocketConnection))

	// Let's get this party started
	cfg := config.ReadConfigs()
	fmt.Println("Server Started [" + cfg.IP + "] on port " + cfg.Port)

	http.ListenAndServe(":"+cfg.Port, nil)
}
