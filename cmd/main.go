package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	info := mux.Vars(r)
	fmt.Print(info["name"])
	fmt.Fprintf(w, "Hello %s\n", info["name"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
