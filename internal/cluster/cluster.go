package cluster

import (
	"strconv"

	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
)

type Cluster struct {
	capacity int
	port     int

	nodes map[string]int
}

func New(capacity int) Cluster {
	return Cluster{
		capacity: capacity,
		port:     11723,
		nodes:    make(map[string]int),
	}
}

func (c *Cluster) CreateNodes() {
	for i := 0; i < c.capacity; i++ {
		address := ":" + strconv.Itoa(c.port)

		c.nodes["localhost"+address] = 0

		c.port++

		go server.New(address)
	}
}
