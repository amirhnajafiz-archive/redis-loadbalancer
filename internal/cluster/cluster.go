package cluster

import (
	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
	"strconv"
)

type Cluster struct {
	capacity int
	port     int

	status map[int]string
}

func New(capacity int) Cluster {
	return Cluster{
		capacity: capacity,
		port:     11723,
		status:   make(map[int]string),
	}
}

func (c *Cluster) Run() {
	for i := 0; i < c.capacity; i++ {
		address := ":" + strconv.Itoa(c.port)

		c.port++

		go server.New(address)
	}
}