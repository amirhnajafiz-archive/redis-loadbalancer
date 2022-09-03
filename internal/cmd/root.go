package cmd

import "github.com/amirhnajafiz/distributed-redis/internal/cmd/server"

func Execute() {
	server.New(":8080")
}
