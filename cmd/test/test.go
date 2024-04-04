package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var clients []*Client
var mu sync.Mutex

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade HTTP connection to WebSocket:", err)
			return
		}

		client := &Client{
			conn: conn,
		}

		mu.Lock()
		clients = append(clients, client)
		mu.Unlock()

		for {
			// Read message from client
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message from WebSocket client:", err)
				break
			}

			// Process message
			processMessage(string(message))
		}

		mu.Lock()
		for i, c := range clients {
			if c == client {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		mu.Unlock()

		conn.Close()
	})

	r.Run(":9090")
}

func processMessage(message string) {
	// Process the message sent from the admin
	// ...

	// Broadcast the message to all connected clients
	mu.Lock()
	defer mu.Unlock()

	for _, client := range clients {
		client.mu.Lock()
		err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
		client.mu.Unlock()
		if err != nil {
			log.Println("Failed to send message to WebSocket client:", err)
		}
	}
}
