package main

import (
	"testing"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

func TestPrefixWrapper(t *testing.T) {
	game_db := NewPrefixWrapper(keyvaluestore.NewRedisStore(), "game_")
	game_db.Set("name", "sdasd")

	if value, _ := game_db.Get("name"); value != "sdasd" {
		t.Fatal("Get failed")
	}
}
