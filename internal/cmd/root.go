package cmd

import "github.com/amirhnajafiz/distributed-redis/internal/cluster"

func Execute() {
	c := cluster.Cluster{
		Capacity: 5,
	}

	c.Register()
}
