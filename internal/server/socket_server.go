package server

import (
	"fmt"
	"go-chat/internal/config"
	"go-chat/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Message represents the structure of a WebSocket message
type Message struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

func handleChatMessage(conn *websocket.Conn, msg Message) {
	// Handle your chat messages here
	// You can switch on msg.Type and perform different actions based on the type
	fmt.Printf("Received chat message: %s\n", msg.Payload)
}

func webSocketHandler(w http.ResponseWriter, r *http.Request, connections map[string]*websocket.Conn, db *database.Database) {
	defer r.Body.Close()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		sendCloseFrame(conn, "Failed to upgrade connection to WebSocket", websocket.CloseInternalServerErr)
		return
	}

	id := r.Header.Get("X-Client-ID")
	if id == "" {
		fmt.Println("Received socket connection but no id received")
		sendCloseFrame(conn, "No client ID provided", websocket.ClosePolicyViolation)
		return
	}

	_, ok := db.User.GetById(id)
	if !ok {
		fmt.Println("No user with this id")
		sendCloseFrame(conn, "User not found", websocket.CloseNormalClosure)
		return
	}

	connections[id] = conn
	fmt.Printf("Client with ID %s connected\n", id)

	// Read and handle incoming messages
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			delete(connections, id) // Remove the connection on error
			sendCloseFrame(conn, "Error reading message", websocket.CloseNormalClosure)
			break
		}

		switch msg.Type {
		case "chat":
			handleChatMessage(conn, msg)
		default:
			fmt.Printf("Unknown message type: %s\n", msg.Type)
		}
	}
}

func SocketServerListen(db *database.Database) {
	var connections = make(map[string]*websocket.Conn)
	port := config.SOCKET_SERVER_PORT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		webSocketHandler(w, r, connections, db)
	})

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	fmt.Printf("Web socket server listening on port %s\n", port)
}

func sendCloseFrame(conn *websocket.Conn, reason string, code int) {
	closeMessage := websocket.FormatCloseMessage(code, reason)
	conn.WriteMessage(websocket.CloseMessage, closeMessage)
	conn.Close()
}
