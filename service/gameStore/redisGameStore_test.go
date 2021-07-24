package main

import "testing"

func TestRedisGameStore(t *testing.T) {
	rdb := NewRedisGameStore()
	if rdb.SetGame("name", "shyam") != nil {
		t.Fatal("Fail")
	}
}
