package main

import (
	"fmt"
	"log"
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

func main() {

	db = keyvaluestore.NewRedisStore()
	game_db = storewrapper.NewPrefixDecorator(db, "game_")
	player_db = storewrapper.NewPrefixDecorator(db, "player_")
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
