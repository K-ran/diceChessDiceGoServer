package gamehandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// checks is the player exists in the db
func checkValidPlayer(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		playerId := params["playerId"]
		_, err := player_db.Get(playerId)
		if err != nil {
			ReturnFailure(rw, "Invalid player id.")
			return
		}
		s.ServeHTTP(rw, r)
	})
}

//checks if the valid game exists in the
func checkValidGame(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		gameId := params["gameId"]
		_, err := game_db.Get(gameId)
		if err != nil {
			ReturnFailure(rw, "Invalid game id.")
			return
		}
		s.ServeHTTP(rw, r)
	})
}
