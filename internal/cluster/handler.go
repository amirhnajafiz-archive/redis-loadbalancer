package cluster

import (
	"encoding/json"
	"net/http"

	"github.com/amirhnajafiz/distributed-redis/internal/http/response"
	"github.com/gin-gonic/gin"
)

const (
	serviceType = "http"
)

// handle
// is the load balancing main request and response handler.
func (c *cluster) handle(ctx *gin.Context) {
	req := ctx.Request
	address := serviceType + "://" + c.getIP() + req.URL.Path

	if req.Method == http.MethodGet {
		resp, err := c.httpClient.Get(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		var responseBody response.PairResponse

		_ = json.NewDecoder(resp.Body).Decode(&responseBody)

		ctx.JSON(resp.StatusCode, responseBody)
	} else if req.Method == http.MethodPost {
		contentTypeHeader := "content-type:" + req.Header.Get("content-type")

		_, err := c.httpClient.Post(address, req.Body, contentTypeHeader)
		if err != nil {
			_ = ctx.Error(err)

			return
		}
	} else if req.Method == http.MethodDelete {
		_, err := c.httpClient.Delete(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}
	}

	ctx.Status(http.StatusNoContent)
}
