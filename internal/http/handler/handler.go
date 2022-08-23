package handler

import (
	"github.com/amirhnajafiz/distributed-redis/internal/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store store.Store
}

func (h *Handler) Insert(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}

func (h *Handler) Get(c *gin.Context) {

}

func (h *Handler) Register(r *gin.RouterGroup) {
	r.POST("/data", h.Insert)
	r.DELETE("/data/{key}", h.Delete)
	r.GET("/data/{key}", h.Get)
}
