package cluster

import (
	"sort"
	"strconv"

	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
	"github.com/amirhnajafiz/distributed-redis/pkg/http_client"
	"github.com/gin-gonic/gin"
)

// LoadBalancingServer
// is the server type.
type LoadBalancingServer interface {
	Health() interface{}
	Start(string string) error
}

// node
// manages to keep information about a service.
type node struct {
	// service ip.
	ip string
	// number of times that the service is used.
	used int
}

// cluster
// manages to monitor clusters and do the load balancing logic.
type cluster struct {
	// Capacity is the maximum number of replicas.
	capacity int
	// HttpClient manages the requests in cluster.
	httpClient *http_client.HTTPClient
	// the array of nodes.
	nodes []*node
}

// create
// manages to create internal services.
func (c *cluster) create() {
	port := 12489

	for i := 0; i < c.capacity; i++ {
		address := ":" + strconv.Itoa(port)
		n := node{
			ip:   "localhost" + address,
			used: 0,
		}

		port++

		c.nodes = append(c.nodes, &n)

		go server.New(address)
	}
}

// getIP
// returns an ip based of the load balancing logic.
func (c *cluster) getIP() string {
	n := c.nodes[0]

	n.used++

	sort.Slice(c.nodes, func(i, j int) bool {
		return c.nodes[i].used < c.nodes[j].used
	})

	return n.ip
}

// New
// creating a new cluster.
func New(capacity int) LoadBalancingServer {
	c := cluster{
		capacity:   capacity,
		httpClient: http_client.New(),
	}

	c.create()

	return &c
}

// Start
// starting the cluster.
func (c *cluster) Start(addr string) error {
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	app.Use(c.handle)

	if err := app.Run(addr); err != nil {
		return err
	}

	return nil
}

// Health
// returns status about cluster.
func (c *cluster) Health() interface{} {
	return c.capacity
}
