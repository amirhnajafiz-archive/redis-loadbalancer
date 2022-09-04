package cluster

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Node struct {
	IP   string
	Used int
}

type Cluster struct {
	Capacity int
	Nodes    []*Node
}

func (c *Cluster) getIP(ctx *gin.Context) {
	n := c.Nodes[0]

	n.Used++

	ctx.String(http.StatusOK, n.IP)
}
