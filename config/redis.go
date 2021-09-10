package config

import (
	"os"
)

type Redis struct {
	CACHE *CacheConfig
}

type CacheConfig struct {
	Host     string
	Port     string
	Db       int
	Ttl      int
	Password string
}

func GetCacheConfig() *Redis {
	return &Redis{
		CACHE: &CacheConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Db:       0,
			Ttl:      600,
			Password: "",
		},
	}
}