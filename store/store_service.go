//go:store_service.go

/*
we are creating  a storage layer for our application
- setting up storage service
- storage api implementation and design
*/

package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// struct wrapper around redis client

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 1 * time.Hour

// initiating the redis server
func InnitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
                Addr:     "localhost:6379", 
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error in redis init %v", err))

	}
	fmt.Printf("Redis server started sucessfully:{%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService

}

func SaveURLMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
	        var message string = fmt.Sprintf("failed saving url | error : %v , shortUrl : %s , originalUrl : %s \n", err, shortUrl, originalUrl)
                log.Print(message) 
	}

}

func RetriveOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
	        var message string = fmt.Sprintf("failed retriving original url | error : %v , shortUrl: %s \n", err, shortUrl)
                log.Print(message) 
	}
	return result
}
