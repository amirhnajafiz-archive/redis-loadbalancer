package cluster

import (
	"strconv"

	"github.com/amirhnajafiz/distributed-redis/internal/cmd/server"
	"github.com/amirhnajafiz/strago"
)

// create
// manages to create internal services.
func create(capacity int) []string {
	var (
		list []string

		port = 12489
	)

	for i := 0; i < capacity; i++ {
		address := ":" + strconv.Itoa(port)
		n := "localhost" + address

		port++

		list = append(list, n)

		go server.New(address)
	}

	return list
}

// New
// creating a new cluster.
func New(capacity, addr int) strago.LoadBalancer {
	nodes := create(capacity)

	opt := strago.NewOptions()

	opt.Port = addr
	opt.Enable = true
	opt.BalancingType = strago.RequestsCount
	opt.Type = "http"

	app := strago.NewServer(opt)

	var services []string

	for _, serv := range nodes {
		services = append(services, serv)
	}

	app.WithServices(services...)

	return app
}
