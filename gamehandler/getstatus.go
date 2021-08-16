package gamehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//response for join api
func GetGetStatusHandler() http.HandlerFunc {
	return allowCORS((checkValidGame(http.HandlerFunc(getStatusHandler))), "*")
}

// handles joining of a game
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId := params["gameId"]
	playerId := params["playerId"]
	w.Header().Add("Content-Type", "application/json")
	var game diceChessGame

	// check if player id and game id matches
	db_game_id, _ := player_db.Get(playerId)
	if db_game_id != gameId {
		ReturnFailure(w, "Trying to access invalid game.")
		return
	}
	gameString, _ := game_db.Get(gameId)
	err := json.Unmarshal([]byte(gameString), &game)
	if err != nil {
		ReturnFailure(w, "Found invalid game data.")
		return
	}
	json.NewEncoder(w).Encode(game)
}
