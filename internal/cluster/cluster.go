package cluster

import (
	"log"
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
	nodes    []*Node
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

func (c *Cluster) handle(ctx *gin.Context) {
	req := ctx.Request
	address := c.getIP() + req.URL.Path

	log.Printf("url: %s\n", address)

	if req.Method == http.MethodGet {
		resp, err := http.Get(address)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(resp.StatusCode, resp.Body)
	} else {
		resp, err := http.Post(address, req.Header.Get("content-type"), req.Body)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(resp.StatusCode, resp.Body)
	}
}

func (c *Cluster) Register() {
	app := gin.Default()

	app.Any("/", c.handle)

	c.create()

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
