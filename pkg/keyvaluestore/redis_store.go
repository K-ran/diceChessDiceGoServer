package keyvaluestore

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// implements key value store interface for redis
type RedisStore struct {
	rdb *redis.Client
}

func (s *RedisStore) Set(key string, value string, ttl int) error {
	err := s.rdb.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
	return err
}

func (s *RedisStore) Get(key string) (string, error) {
	return s.rdb.Get(ctx, key).Result()
}

func (s *RedisStore) Update(key string, value string, ttl int) (string, error) {
	if _, err := s.Get(key); err == nil {
		s.Set(key, value, ttl)
		return value, nil
	} else {
		return "", &KeyError{"Key not found"}
	}
}

func (s *RedisStore) Delete(key string) {
	s.rdb.Del(ctx, key)
}

func (s *RedisStore) Connect() {
	s.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if s.rdb == nil {
		panic("Not able to connect to Redis")
	} else {
		log.Printf("Connect to Redis.\n")
	}
}

func (s *RedisStore) Disconnect() {
	s.rdb.Close()
}

//Returns instance of redis game store
func NewRedisStore() *RedisStore {
	instance := new(RedisStore)
	instance.Connect()
	return instance
}
