package handler

import (
	"net/http"
	"time"

	"github.com/amirhnajafiz/distributed-redis/internal/http/request"
	"github.com/amirhnajafiz/distributed-redis/internal/http/response"
	"github.com/amirhnajafiz/distributed-redis/internal/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store store.Store
}

func (h *Handler) insert(c *gin.Context) {
	var pair request.NewPairRequest

	if err := c.BindJSON(&pair); err != nil {
		_ = c.Error(err)

		return
	}

	if err := h.Store.Put(pair.Key, pair.Value); err != nil {
		_ = c.Error(err)

		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) delete(c *gin.Context) {
	err := h.Store.Trash(c.Param("key"))
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) get(c *gin.Context) {
	value, err := h.Store.Pull(c.Param("key"))
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, response.PairResponse{
		Key:   c.Param("key"),
		Value: value,
		Date:  time.Now(),
	})
}

func (h *Handler) Register(r *gin.RouterGroup) {
	r.POST("/data", h.insert)
	r.DELETE("/data/{key}", h.delete)
	r.GET("/data/{key}", h.get)
}
