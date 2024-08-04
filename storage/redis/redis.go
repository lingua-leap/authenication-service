package redisr

import (
	red "github.com/redis/go-redis/v9"
)

func ConnectRedis() *red.Client {
	client := red.NewClient(&red.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
