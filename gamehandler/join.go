package gamehandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/k-ran/diceChessDiceServer/middleware"

	"github.com/gorilla/mux"
	"github.com/k-ran/diceChessDiceServer/internal/storewrapper"
)

//response for join api
type joinResponse struct {
	RespType  int           `json:"type"` //type of response
	PlayerId  string        `json:"playerId"`
	GameState diceChessGame `json:"gameState"`
}

func GetJoinHandler() http.HandlerFunc {
	return allowCORS(middleware.InputJoinCheck(http.HandlerFunc(joinHandler)), "*")
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
		return
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
	playerId := storewrapper.GreateUniquePlayerId(player_db, gameId)

	value, err := json.Marshal(game)
	if err != nil {
		log.Println("Failed to parse inputs")
		ReturnFailure(w, "Failed to parse inputs.")
		return
	}

	//save the game state
	game_db.Set(gameId, string(value), 0)

	//return response
	resp := joinResponse{RESPONSE_JOIN, playerId, game}
	json.NewEncoder(w).Encode(resp)
}
