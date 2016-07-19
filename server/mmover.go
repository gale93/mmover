package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

type config struct {
	Port string `json:"port"`
}

func main() {

	// Reading configs from file
	var cfg config
	file, err := ioutil.ReadFile("config.cfg")

	if err != nil {
		fmt.Println("Error Reading the config file\n" + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(file, &cfg)

	if err != nil {
		fmt.Println("Error decoding the config file\n" + err.Error())
		os.Exit(1)
	}

	// Redirecting requests to resources' folder
	http.Handle("/", http.FileServer(http.Dir("../client")))

	// data needed from the device app
	http.HandleFunc("/data", get)

	// Activate ws handler to retrieve browser data
	http.Handle("/socket", websocket.Handler(connection))

	// Let's get this party started
	fmt.Println("Server Started on port " + cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}
