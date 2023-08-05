package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"shorturl/packages/utilities"
	"time"
)

type RedisSvc struct {
	ctx context.Context
	rd  *redis.Client
}

type RedisSvcConfig struct {
	Password string
	Host     string
	Port     int32
	Db       int32
}

type RedisSetOptions struct {
	expire int
}

func NewRedisSvc(config RedisSvcConfig) *RedisSvc {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     utilities.ConcatStrings(config.Host, ":", string(config.Port)),
		Password: config.Password, // no password set
		DB:       int(config.Db),  // use default DB
	})

	if pong, error := client.Ping(ctx).Result(); error != nil {
		panic(error)
	} else {
		fmt.Println(pong)
	}

	return &RedisSvc{rd: client, ctx: ctx}
}

func (r RedisSvc) SetAsync(key string, value string, result chan interface{}, options *RedisSetOptions) {
	status := r.rd.Set(r.ctx, key, value, 0)
	if status.Err() != nil {
		panic(status)
	}
	if options.expire != 0 {
		r.Expire(key, int32(options.expire))
	}

	if response, err := status.Result(); err != nil {
		panic(err)
	} else {
		result <- response
	}
}

func (r RedisSvc) Put(namespace, key, value string, options *RedisSetOptions) bool {
	namespace = utilities.ConcatStrings(namespace, "_", "shorturls")
	if response, _error := r.rd.HSetNX(r.ctx, namespace, key, value).Result(); _error != nil {
		panic(_error)
	} else {
		r.Expire(namespace, int32(options.expire))
		return response
	}
}

func (r RedisSvc) Get(namespace, key string) interface{} {
	namespace = utilities.ConcatStrings(namespace, "_", "shorturls")
	if response, _error := r.rd.HGet(r.ctx, namespace, key).Result(); _error != nil {
		panic(_error)
	} else {
		return response
	}
}

func (r RedisSvc) GetAsync(namespace, key string, response chan interface{}) {
	response <- r.Get(namespace, key)
}

func (r RedisSvc) PutAsync(namespace, key, value string, result chan interface{}) {
	response := r.Put(namespace, key, value, nil)

	result <- response
}

func (r RedisSvc) Expire(key string, seconds int32) {
	if err := r.rd.Expire(r.ctx, key, time.Second*time.Duration(seconds)).Err(); err != nil {
		panic(err)
	}
}

// https://tutorialedge.net/courses/go-data-structures-course/
