package storewrapper

import (
	"testing"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

func TestPrefixWrapper(t *testing.T) {
	game_db := NewPrefixDecorator(keyvaluestore.NewRedisStore(), "game_")
	game_db.Set("name", "sdasd", 10)

	if value, _ := game_db.Get("name"); value != "sdasd" {
		t.Fatal("Get failed")
	}
}
