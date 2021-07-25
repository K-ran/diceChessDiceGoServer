package gamehandler

import (
	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"

	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
)

var db keyvaluestore.KeyValueStore
var game_db keyvaluestore.KeyValueStore
var player_db keyvaluestore.KeyValueStore

const DB_ENTRY_TLL int = 900 //15minutes

func InitDb() {
	db = keyvaluestore.NewRedisStore()
	game_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "game_"), DB_ENTRY_TLL)

	player_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "player_"), DB_ENTRY_TLL)
}
