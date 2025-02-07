package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type RedisManager struct {
	client *redis.Client
	ctx    context.Context
	once   sync.Once
}

var redisOnce sync.Once
var redisInstance *RedisManager

func GetRedisManager(addr, password string, db int) *RedisManager {
	redisOnce.Do(
		func() {
			log.Println("Initializing redis connection...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			client := redis.NewClient(
				&redis.Options{
					Addr:     addr,
					Password: password,
					DB:       db,
				},
			)

			if err := client.Ping(ctx).Err(); err != nil {
				log.Fatalf("Could not connect to Redis: %v", err)
			}

			redisInstance = &RedisManager{
				client: client,
				ctx:    context.Background(),
			}
		},
	)

	return redisInstance
}

func (r *RedisManager) Close() {
	if err := r.client.Close(); err != nil {
		log.Fatalf("Could not close Redis: %v", err)
	}
}
