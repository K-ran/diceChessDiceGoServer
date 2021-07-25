package KeyValueStore

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// implements key value store interface for redis
type RedisStore struct {
	rdb *redis.Client
}

func (s *RedisStore) Set(key string, value string) error {
	err := s.rdb.Set(ctx, key, value, 0).Err()
	return err
}

func (s *RedisStore) Get(key string) (string, error) {

	return s.rdb.Get(ctx, key).Result()
}

func (s *RedisStore) Update(key string, value string) (string, error) {
	if _, err := s.Get(key); err == nil {
		s.Set(key, value)
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
		panic("Cannot connect to db")
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
