package cmd

import (
	"github.com/amirhnajafiz/distributed-redis/internal/cluster"
	"github.com/amirhnajafiz/distributed-redis/pkg/http_client"
)

func Execute() {
	c := cluster.Cluster{
		Capacity:   5,
		HttpClient: http_client.New(),
	}

	c.Register()
}
