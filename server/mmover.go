package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type data struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`

	UpdateRatems int `json:"update_rate_ms"`
}

var app data

func get(w http.ResponseWriter, r *http.Request) {
	dt, _ := json.Marshal(app)
	w.Write(dt)
}

func connection(ws *websocket.Conn) {
	defer ws.Close()
	for {
		// Get message from client
		if err := websocket.JSON.Receive(ws, &app); err != nil {
			log.Println("A socket has been disconnected.")
			return
		}
	}
}

func main() {

	// Redirecting requests to resources' folder
	http.Handle("/", http.FileServer(http.Dir("../client")))

	// data needed from the device app
	http.HandleFunc("/data", get)

	// Activate ws handler to retrieve browser data
	http.Handle("/socket", websocket.Handler(connection))

	// Let's get this party started
	http.ListenAndServe(":9393", nil)
}
