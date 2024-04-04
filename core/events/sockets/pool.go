package sockets

import (
	"flick_tickets/common/log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan interface{}
	upgrader  websocket.Upgrader
	mutex     sync.Mutex
	rooms     map[string]map[*websocket.Conn]bool
}

func NewServer() *server {
	server := &server{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan interface{}),
		rooms:     make(map[string]map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return server
}

func (s *server) runSocket(c *gin.Context, room string) {

	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Infof("error websocket : ", err)
		return
	}
	s.mutex.Lock()

	if _, ok := s.rooms[room]; !ok {
		s.rooms[room] = make(map[*websocket.Conn]bool)
	}

	s.rooms[room][conn] = true
	s.clients[conn] = true
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.clients, conn)
		if len(s.rooms[room]) == 0 {
			delete(s.rooms, room)
		}
		s.mutex.Unlock()
		conn.Close()
	}()

	log.Infof("WebSocket connection established") // Log khi kết nối thành công

	go s.handleMessages(room)
	for {
		var message interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Error(err, "close websocket error ! ")
			break
		}
		s.broadcast <- message
		//s.writeMessage(conn, message)
	}
}

func (s *server) handleMessages(room string) {

	for message := range s.broadcast {
		s.mutex.Lock()
		client, ok := s.rooms[room]
		if !ok {
			s.mutex.Unlock()
			continue
		}
		for client := range client {
			err := client.WriteJSON(message)
			if err != nil {
				log.Errorf(err, "Failed to write message to WebSocket client: %s")
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mutex.Unlock()
	}

}
func (s *server) writeMessage(conn *websocket.Conn, message interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := conn.WriteJSON(message)
	if err != nil {
		log.Infof("Failed to write message to WebSocket client: %v", err)
		conn.Close()
		delete(s.clients, conn)
		return err
	}

	return nil
}
