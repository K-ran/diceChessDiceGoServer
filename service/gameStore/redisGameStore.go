package main

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisGameStore struct {
	rdb *redis.Client
}

func (s *RedisGameStore) SetGame(key string, value string) error {
	err := s.rdb.Set(ctx, key, value, 0).Err()
	return err
}

func (s *RedisGameStore) GetGame(key string, value string) (string, error) {
	return "", nil
}

func (s *RedisGameStore) UpdateGame(key string, value string) (string, error) {
	return "", nil
}

func (s *RedisGameStore) DeleteGame(key string) {

}

func (s *RedisGameStore) Connect() {
	s.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if s.rdb == nil {
		panic("Cannot connect to db")
	}
}

var instance *RedisGameStore
var lock = &sync.Mutex{}

//Returns Singleton instance of redis game store
func NewRedisGameStore() *RedisGameStore {
	lock.Lock()
	defer lock.Unlock()

	if instance != nil {
		return instance
	} else {
		instance = new(RedisGameStore)
		instance.Connect()
		return instance
	}
}
