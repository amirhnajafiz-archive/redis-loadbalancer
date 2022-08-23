package store

import "github.com/go-redis/redis/v9"

// Connect method makes a connection to
// redis client.
func Connect(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
}
