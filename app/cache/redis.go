package cache

import (
	"fmt"

	"skeleton_project/config"

	"github.com/go-redis/redis"
)

type Cache struct{
	Client *redis.Client
	Config *config.Redis
}

func (c *Cache) InitClient(config *config.Redis) {
	Addr := fmt.Sprintf("%s:%s", config.CACHE.Host, config.CACHE.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: config.CACHE.Password,
		DB:       config.CACHE.Db,
	})

	c.Client = client
	c.Config = config

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
}

func (c *Cache) Get(key string) (string, error) {
	return c.Client.Get(key).Result()
}

func (c *Cache) Set(key string, value interface{}) error {
	return c.Client.Set(key, value, 0).Err()
}

func (c *Cache) Delete(key string) *redis.IntCmd {
	return c.Client.Del(key)
}
