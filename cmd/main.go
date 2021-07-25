package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	info := mux.Vars(r)
	fmt.Print(info["name"])
	fmt.Fprintf(w, "Hello %s\n", info["name"])
}

func main() {
	var inst keyvaluestore.KeyValueStore
	inst = keyvaluestore.NewRedisStore()
	inst.Set("name", "ram_ram")
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
