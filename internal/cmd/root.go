package cmd

import (
	"github.com/amirhnajafiz/distributed-redis/internal/http/handler"
	"github.com/amirhnajafiz/distributed-redis/internal/store"
	"github.com/gin-gonic/gin"
)

func Execute() {
	app := gin.Default()

	h := handler.Handler{
		Store: store.Store{
			Conn: store.Connect("localhost:6379"),
		},
	}

	h.Register(app.Group("/api"))

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
