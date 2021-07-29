package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	r.HandleFunc("/create/{playerName}/{gameName}/{dieNum}", gamehandler.GetCreateHttpHandler()).Methods("POST")
	r.HandleFunc("/join/{playerName}/{gameId}", gamehandler.GetJoinHandler()).Methods("POST")
	r.HandleFunc("/getstatus/{playerId}/{gameId}", gamehandler.GetGetStatusHandler()).Methods("GET")
	http.Handle("/", r)

	address := os.Getenv("GO_PORT")
	if address == "" {
		address = "8081"
	}

	log.Println("Server starting at port " + address)
	log.Fatal(http.ListenAndServe(":"+address, nil))
}
