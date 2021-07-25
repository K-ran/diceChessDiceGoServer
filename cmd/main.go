package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/gamehandler"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	info := mux.Vars(r)
	fmt.Print(info["name"])
	fmt.Fprintf(w, "Hello %s\n", info["name"])
}

func main() {
	gamehandler.InitDb()
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomeHandler)
	http.Handle("/", r)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}
