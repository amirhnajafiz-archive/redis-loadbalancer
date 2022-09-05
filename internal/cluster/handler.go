package cluster

import (
	"encoding/json"
	"net/http"

	"github.com/amirhnajafiz/distributed-redis/internal/http/response"
	"github.com/gin-gonic/gin"
)

func (c *Cluster) handle(ctx *gin.Context) {
	req := ctx.Request
	address := "http://" + c.getIP() + req.URL.Path

	if req.Method == http.MethodGet {
		resp, err := c.HttpClient.Get(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		var responseBody response.PairResponse

		_ = json.NewDecoder(resp.Body).Decode(&responseBody)

		ctx.JSON(resp.StatusCode, responseBody)
	} else if req.Method == http.MethodPost {
		_, err := c.HttpClient.Post(address, req.Body, req.Header.Get("content-type"))
		if err != nil {
			_ = ctx.Error(err)

			return
		}
	} else if req.Method == http.MethodDelete {
		_, err := c.HttpClient.Delete(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}
	}

	ctx.Status(http.StatusNoContent)
}
