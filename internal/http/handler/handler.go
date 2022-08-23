package handler

import (
	"net/http"
	"time"

	"github.com/amirhnajafiz/distributed-redis/internal/http/response"
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
	r.POST("/data", h.Insert)
	r.DELETE("/data/{key}", h.Delete)
	r.GET("/data/{key}", h.Get)
}
