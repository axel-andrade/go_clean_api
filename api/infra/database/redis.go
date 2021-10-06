package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func ConnectRedisDB() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr:     dsn,    //redis port
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	var ctx = context.TODO()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisDB() *redis.Client {
	return client
}
