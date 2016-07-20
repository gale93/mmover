package connection

import (
	"encoding/json"
	"log"
	"mmover/server/config"
	"net/http"

	"golang.org/x/net/websocket"
)

//GetData Return all data in a json array needed for device
func GetData(w http.ResponseWriter, r *http.Request) {
	dt, _ := json.Marshal(app)
	w.Write(dt)
}

// WebSocketConnection is the connection used from the browser
func WebSocketConnection(ws *websocket.Conn) {
	defer ws.Close()

	// send to new socket the starting position
	msg := "{\"header\": \"starting_pos\", \"body\": { \"lat\": \"" + config.Cfg.StartingLat + "\", \"lng\": \"" + config.Cfg.StartingLng + "\"}}"
	websocket.Message.Send(ws, msg)

	for {
		// Get message from client
		if err := websocket.JSON.Receive(ws, &app); err != nil {
			log.Println("A socket has been disconnected.")
			return
		}
	}
}
