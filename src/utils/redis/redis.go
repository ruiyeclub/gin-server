package redis

import (
	"encoding/json"
	"fmt"
	"gin-server/src/utils/config"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var Server *redisServer

type redisServer struct {
	rs *redis.Client
}

func NewRedisServer() *redisServer {
	r := new(redisServer)
	r.loadRedisConfig()
	return r
}

func init() {
	Server = NewRedisServer()
}

func (r *redisServer) loadRedisConfig() {

	if r.rs != nil {
		return
	}

	r_host := config.Config.Redis.Host
	r_port := config.Config.Redis.Port
	r_host = fmt.Sprintf("%s:%s", r_host, r_port)

	r_password := config.Config.Redis.Password
	r_db := config.Config.Redis.Database

	client := redis.NewClient(&redis.Options{
		Addr:       r_host,
		Password:   r_password,
		DB:         r_db,
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Println("Redis初始化失败 : ", err)
	} else {
		log.Println("Redis初始化成功")
	}
	r.rs = client
}

func (r *redisServer) RC() *redis.Client {
	return r.rs
}

func (r *redisServer) Set(key string, value interface{}, expiration time.Duration) {
	val := r.rs.Set(key, value, expiration).Val()
	log.Println("redis set result: ", val)
}

func (r *redisServer) Get(key string) string {
	result := r.rs.Get(key).Val()
	return result
}

func (r *redisServer) GetByAny(key string, any interface{}) error {
	return r.rs.Get(key).Scan(any)
}

func (r *redisServer) GetStruct(key string, any interface{}) {
	resultStr := r.Get(key)
	if resultStr != "" {
		json.Unmarshal([]byte(resultStr), any)
	}
}
