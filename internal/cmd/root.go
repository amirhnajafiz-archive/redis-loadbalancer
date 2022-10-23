package cmd

import (
	"github.com/amirhnajafiz/distributed-redis/internal/cluster"
)

func Execute() {
	c := cluster.New(3, 8080)

	if err := c.Start(); err != nil {
		panic(err)
	}
}
