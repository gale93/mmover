package connection

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	dt, _ := json.Marshal(app)
	w.Write(dt)
}

func WebSocketConnection(ws *websocket.Conn) {
	defer ws.Close()

	for {
		// Get message from client
		if err := websocket.JSON.Receive(ws, &app); err != nil {
			log.Println("A socket has been disconnected.")
			return
		}
	}
}
