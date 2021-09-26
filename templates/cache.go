package templates

var CacheProvider = `
package cache

import "time"

type CacheProvider interface {
	Get(key string) (result []byte, err error)
	Set(key string, value string, expire time.Duration) (err error)
	Delete(key string) (err error)
}
`

var RedisProvider = `
package cache

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

/* CHANGE REDIS_HOST and REDIS_PREFIX */

type redisCacheProvider struct {
	cacheClient *redis.Client
}

func NewRedisProvider() CacheProvider {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "REDIS_HOST",
		DB:   0,
	})
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connected ok.")
	return &redisCacheProvider{
		cacheClient: redisClient,
	}
}

func (r *redisCacheProvider) Get(key string) (result []byte, err error) {
	return r.cacheClient.Get(context.Background(), r.addPrefix(key)).Bytes()
}

func (r *redisCacheProvider) Set(key string, value string, expire time.Duration) (err error) {
	return r.cacheClient.Set(context.Background(), r.addPrefix(key), value, expire).Err()
}

func (r *redisCacheProvider) Delete(key string) (err error) {
	return r.cacheClient.Del(context.Background(), r.addPrefix(key)).Err()
}

func (r *redisCacheProvider) addPrefix(key string) string {
	return fmt.Sprintf("%s%S", "REDIS_PREFIX", key)
}
`