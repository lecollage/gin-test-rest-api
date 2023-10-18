package resources

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var Redis *redis.Client

func RedisConnect() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	dbRaw := os.Getenv("REDIS_DB")

	addr := fmt.Sprintf("%s:%s", host, port)
	db, _ := strconv.Atoi(dbRaw)

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
