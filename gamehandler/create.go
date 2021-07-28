package gamehandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
	"github.com/k-ran/diceChessDiceServer/middleware"
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

// response for create api
type createResponse struct {
	Status     string `json:"status"`
	GameId     string `json:"gameId"`
	PlayerId   string `json:"playerId"`
	GameName   string `json:"gameName"`
	PlayerName string `json:"playerName"`
	State      int    `json:"state"`
	DieNum     string `json:"dieNum"`
}

// Generate the Create the http handler and return
func GetCreateHttpHandler() http.HandlerFunc {
	return middleware.InputCreateCheck(
		http.HandlerFunc(createHandler))
}

// Actual create handler function
func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateHandler Called...\n")
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	gameName := params["gameName"]
	playerName := params["playerName"]
	dieNum, _ := strconv.Atoi(params["dieNum"])
	playerId := storewrapper.GreateUniquePlayerId(player_db)
	gameId := storewrapper.GreateUniqueGameId(game_db)
	game := NewDiceChessGame(gameId, gameName, playerName, "", dieNum)
	value, err := json.Marshal(game)
	if err != nil {
		log.Println("Failed to parse inputs")
		ReturnFailure(w, "Failed to parse inputs")
		return
	}
	game_db.Set(gameId, string(value), 0)
	resp := createResponse{"0", gameId, playerId, gameName, playerName, int(WAITING), strconv.Itoa(dieNum)}
	json.NewEncoder(w).Encode(resp)
}
