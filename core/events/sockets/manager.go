package sockets

import "github.com/gin-gonic/gin"

type ManagerClient struct {
	sockets map[string]*server
}

func NewManagerClient() *ManagerClient {
	return &ManagerClient{
		sockets: make(map[string]*server),
	}
}
func (mc *ManagerClient) ServerWs(c *gin.Context) {

	id := c.Query("id")

	if _, ok := mc.sockets[id]; !ok {
		mc.sockets[id] = NewServer() // init room
	}
	mc.sockets[id].runSocket(c, id)
}
