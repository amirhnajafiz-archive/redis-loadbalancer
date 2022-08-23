package handler

import "github.com/amirhnajafiz/distributed-redis/internal/store"

type Handler struct {
	Store store.Store
}
