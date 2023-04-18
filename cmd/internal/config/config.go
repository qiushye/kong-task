package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mongo MongoConfig
	Redis cache.CacheConf
}

type MongoConfig struct {
	Url string
	DB  string `json:",default=kong"`
}
