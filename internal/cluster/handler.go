package cluster

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amirhnajafiz/distributed-redis/internal/http/response"
	"github.com/gin-gonic/gin"
)

func (c *Cluster) handle(ctx *gin.Context) {
	req := ctx.Request
	address := "http://" + c.getIP() + req.URL.Path

	log.Printf("url: %s\n", address)

	if req.Method == http.MethodGet {
		resp, err := http.Get(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		var responseBody response.PairResponse

		_ = json.NewDecoder(resp.Body).Decode(&responseBody)

		ctx.JSON(resp.StatusCode, responseBody)
	} else {
		resp, err := http.Post(address, req.Header.Get("content-type"), req.Body)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(resp.StatusCode, resp.Body)
	}

	ctx.Status(http.StatusNoContent)
}
