package gamehandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/k-ran/diceChessDiceServer/middleware"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
)

//response for join api
type joinResponse struct {
	Status     string `json:"status"`
	GameId     string `json:"gameId"`
	PlayerId   string `json:"playerId"`
	GameName   string `json:"gameName"`
	PlayerName string `json:"playerName"`
	DieNum     string `json:"dieNum"`
}

func GetJoinHandler() http.HandlerFunc {
	return middleware.InputJoinCheck(http.HandlerFunc(joinHandler))
}

// handles joining of a game
func joinHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId := params["gameId"]
	playerName := params["playerName"]
	w.Header().Add("Content-Type", "application/json")
	var game diceChessGame
	gameString, _ := game_db.Get(gameId)
	err := json.Unmarshal([]byte(gameString), &game)
	if err != nil {
		ReturnFailure(w, "Found invalid game data.")
	}

	//populate the missing data in the game structure
	game.Player2 = player{playerName}
	if game.State == RUNNING {
		ReturnFailure(w, "Game room full!")
		return
	} else {
		game.State = RUNNING
	}

	//generate new player id
	playerId := storewrapper.GreateUniquePlayerId(player_db)

	value, err := json.Marshal(game)
	if err != nil {
		log.Println("Failed to parse inputs")
		ReturnFailure(w, "Failed to parse inputs")
		return
	}

	//save the game state
	game_db.Set(gameId, string(value), 0)

	//return response
	resp := joinResponse{"0", gameId, playerId, game.GameName, playerName, strconv.Itoa(len(game.Dice))}
	json.NewEncoder(w).Encode(resp)
}
