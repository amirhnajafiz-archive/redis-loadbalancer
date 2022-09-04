<h1 align="center">
Distributed Redis
</h1>

Creating a distributed Redis client with Golang and Clustering. The idea is to create a clustering system for Redis clients, maybe a like a load balancer.
Therefore, we can speed up the process of handling clients.

## How to use the project?
Clone the repository and start the balancing server:
```shell
go run main.go
```

Now you can get a URL for redis connection by:
```shell
curl lcoalhost:8080/
```

Each time you request for a cluster, you will get a new IP.

## Test
You can set a redis cluster with docker compose:
```shell
docker-compose up -d
```