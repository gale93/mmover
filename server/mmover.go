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
	config.ReadConfigs()
	fmt.Println("mmover [" + config.Cfg.IP + "] on port [" + config.Cfg.Port + "]")
	fmt.Println("You can start using it under one of these links:")
	fmt.Println("=>\thttp://localhost:" + config.Cfg.Port + "\n=>\thttp://" + config.Cfg.IP + ":" + config.Cfg.Port)

	http.ListenAndServe(":"+config.Cfg.Port, nil)
}
