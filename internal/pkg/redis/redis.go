package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCfg struct {
	//addr
	Addr string `ini:"Addr"`
	//password
	Password string `ini:"Password"`
	//db
	DB int `ini:"DB"`
	//pool size
	PoolSize int `ini:"PoolSize"`
}

type Redis interface {
	Engine() *redis.Client
}

type redisImpl struct {
	rdb *redis.Client
	cfg *RedisCfg
}

func connect(cfg *RedisCfg) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return rdb
}

func NewRedis(cfg *RedisCfg) Redis {
	return &redisImpl{rdb: connect(cfg), cfg: cfg}
}

func (r *redisImpl) Engine() *redis.Client {
	return r.rdb
}
