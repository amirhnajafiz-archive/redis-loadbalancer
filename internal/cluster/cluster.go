package cluster

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
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

func (c *Cluster) create() {
	port := 12489

	for i := 0; i < c.Capacity; i++ {
		address := ":" + strconv.Itoa(port)
		n := Node{
			IP:   "localhost" + address,
			Used: 0,
		}

		port++

		c.Nodes = append(c.Nodes, &n)

		go server.New(address)
	}
}

func (c *Cluster) getIP(ctx *gin.Context) {
	n := c.Nodes[0]

	n.Used++

	sort.Slice(c.Nodes, func(i, j int) bool {
		return c.Nodes[i].Used < c.Nodes[j].Used
	})

	ctx.String(http.StatusOK, n.IP)
}

func (c *Cluster) Register() {

}
