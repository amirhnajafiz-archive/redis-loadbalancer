package cmd

import (
	"github.com/amirhnajafiz/distributed-redis/internal/cluster"
)

func Execute() {
	c := cluster.New(3)

	if err := c.Start(":8080"); err != nil {
		panic(err)
	}
}
