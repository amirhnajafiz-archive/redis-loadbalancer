package cluster

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
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

	sort.Slice(c.Nodes, func(i, j int) bool {
		return c.Nodes[i].Used < c.Nodes[j].Used
	})

	ctx.String(http.StatusOK, n.IP)
}
