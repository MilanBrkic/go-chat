package server

import (
	"fmt"
	"go-chat/internal/config"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketHandler(w http.ResponseWriter, r *http.Request, connections map[string]*websocket.Conn) {
	defer r.Body.Close()

	id := r.Header.Get("X-Client-ID")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	connections[id] = conn
	fmt.Printf("Client with ID %s connected\n", id)
}

func SocketServerListen() {
	var connections = make(map[string]*websocket.Conn)
	port := config.SOCKET_SERVER_PORT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		webSocketHandler(w, r, connections)
	})

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	fmt.Printf("Web socket server listening on port %s\n", port)
}
