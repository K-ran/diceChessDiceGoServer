package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	info := mux.Vars(r)
	fmt.Print(info["name"])
	fmt.Fprintf(w, "Hello %s\n", info["name"])
}

var db keyvaluestore.KeyValueStore
var game_db keyvaluestore.KeyValueStore
var player_db keyvaluestore.KeyValueStore

const DB_ENTRY_TLL int = 900 //15minutes

func main() {

	db = keyvaluestore.NewRedisStore()
	game_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "game_"), DB_ENTRY_TLL)

	player_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "player_"), DB_ENTRY_TLL)

	game_db.Set("NAME", "KARAN", 0)
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomeHandler)
	http.Handle("/", r)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}
