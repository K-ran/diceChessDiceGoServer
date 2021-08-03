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

	r.HandleFunc("/create/{playerName}/{gameName}/{dieNum}", gamehandler.GetCreateHttpHandler()).Methods("POST", "OPTIONS")
	r.HandleFunc("/join/{gameId}/{playerName}", gamehandler.GetJoinHandler()).Methods("POST")
	r.HandleFunc("/getstatus/{gameId}/{playerId}", gamehandler.GetGetStatusHandler()).Methods("GET")
	r.HandleFunc("/roll/{gameId}/{playerId}", gamehandler.GetRollHandler()).Methods("GET")
	http.Handle("/", r)

	address := os.Getenv("GO_PORT")
	if address == "" {
		address = "8081"
	}

	log.Println("Server starting at port " + address)
	log.Fatal(http.ListenAndServe(":"+address, nil))
}
