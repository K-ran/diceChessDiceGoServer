package gamehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//response for join api
type rollResponse struct {
	RespType  int          `json:"type"` //type of response
	Dies      []diceStruct `json:"dice"`
	RollCount int          `json:"rollCount"`
}

func GetRollHandler() http.HandlerFunc {
	return allowCORS((checkValidGame(http.HandlerFunc(rollHandler))), "*")
}

// handles joining of a game
func rollHandler(w http.ResponseWriter, r *http.Request) {
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

	for i := 0; i < len(game.Dice); i++ {
		game.Dice[i].Rollit()
	}

	//increase the roll count
	game.RollCount++

	value, _ := json.Marshal(game)
	//save the game state
	game_db.Set(gameId, string(value), 0)
	rollResp := rollResponse{RESPONSE_ROLL, game.Dice, game.RollCount}
	json.NewEncoder(w).Encode(rollResp)
}
