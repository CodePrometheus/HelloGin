package tool

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore RedisStore

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	RediStore = RedisStore{client: client}
	return &RediStore
}

// set
func (rs RedisStore) Set(id string, value string) {
	err := rs.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// get
func (rs RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}

// verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
