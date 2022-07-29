package cache

import (
	"context"
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"    // package used to read the .env file

	"github.com/go-redis/redis/v8"
	
)

var rdb *redis.Client = nil
var ctx = context.Background()

func GetRedis() (*redis.Client, context.Context) {
	if rdb != nil {
		return rdb, ctx
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis Init Failed")
	}
	return rdb, ctx
}

func SetValue(key, value string, time time.Duration) error {
	rdb, ctx := GetRedis()
	err := rdb.Set(ctx, key, value, time).Err()
	return err
}

func GetValue(key string) (string, error) {
	rdb, ctx := GetRedis()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func DeleteValue(key string) error {
	rdb, ctx := GetRedis()
	err := rdb.Del(ctx, key).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}