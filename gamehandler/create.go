package gamehandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
)

var db keyvaluestore.KeyValueStore
var game_db keyvaluestore.KeyValueStore
var player_db keyvaluestore.KeyValueStore

const DB_ENTRY_TLL int = 900 //15minutes

// Initilize the databases
func InitDb() {
	db = keyvaluestore.NewRedisStore()
	game_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "game_"), DB_ENTRY_TLL)

	player_db = storewrapper.NewTtlDeocrator(
		storewrapper.NewPrefixDecorator(
			db, "player_"), DB_ENTRY_TLL)
}

type createResponse struct {
	Status     int    `json:"status"`
	GameId     string `json:"gameId"`
	PlayerId   string `json:"playerId"`
	GameName   string `json:"gameName"`
	PlayerName string `json:"playerName"`
	State      int    `json:"state"`
	DieNum     int    `json:"dieNum"`
}

// Generate the Create the http handler and return
func GetCreateHttpHandler() http.HandlerFunc {
	return inputCreateCheck(http.HandlerFunc(CreateHandler))
}

// Actual create handler function
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateHandler Called...\n")
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	gameName := params["gameName"]
	playerName := params["playerName"]
	dieNum, _ := strconv.Atoi(params["dieNum"])
	playerId := storewrapper.GreateUniquePlayerId(player_db)
	gameId := storewrapper.GreateUniquePlayerId(game_db)
	resp := createResponse{0, gameId, playerId, gameName, playerName, int(WAITING), dieNum}
	json.NewEncoder(w).Encode(resp)
	// w.Write([]byte("hi"))
}
