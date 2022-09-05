package server

import (
	"github.com/amirhnajafiz/distributed-redis/internal/http/handler"
	"github.com/amirhnajafiz/distributed-redis/internal/store"
	"github.com/gin-gonic/gin"
)

func New(address string) {
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	h := handler.Handler{
		Store: store.Store{
			Conn: store.Connect("localhost:6379"),
		},
	}

	h.Register(app.Group("/api"))

	if err := app.Run(address); err != nil {
		panic(err)
	}
}
