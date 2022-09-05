package cluster

import (
	"sort"
	"strconv"

	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
	"github.com/amirhnajafiz/distributed-redis/pkg/http_client"
	"github.com/gin-gonic/gin"
)

type Node struct {
	IP   string
	Used int
}

type Cluster struct {
	Capacity   int
	HttpClient *http_client.HTTPClient

	nodes []*Node
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

		c.nodes = append(c.nodes, &n)

		go server.New(address)
	}
}

func (c *Cluster) getIP() string {
	n := c.nodes[0]

	n.Used++

	sort.Slice(c.nodes, func(i, j int) bool {
		return c.nodes[i].Used < c.nodes[j].Used
	})

	return n.IP
}

func (c *Cluster) Register() {
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	app.Use(c.handle)

	c.create()

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
