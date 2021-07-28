package gamehandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const GAME_NAME_LENGTH = 128
const PLAYER_NAME_LENGTH = 128
const MAX_DIE_VALUE = 3

//Checks string for aplhanumeric characters
func isAlphaNumeric(str string) bool {
	for _, character := range str {
		if (character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}

//Checks string for alphabets only
func isAlpha(str string) bool {
	for _, character := range str {
		if (character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') {
			continue
		} else {
			return false
		}
	}
	return true
}

// Middleware to check the input for create api
func inputCreateCheck(s http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		foundError := 0
		errorString := ""
		params := mux.Vars(r)
		gameName := params["gameName"]
		playerName := params["playerName"]
		dieNum, err := strconv.Atoi(params["dieNum"])

		if err != nil {
			foundError++
			errorString += "Die number should be between 1-3."
		} else {
			if dieNum <= 0 || dieNum > MAX_DIE_VALUE {
				foundError++
				errorString += "Die number should be between 1-3."
			}
		}

		if len(gameName) > GAME_NAME_LENGTH || len(gameName) <= 0 {
			foundError++
			errorString += "Invalid game name length."
		}
		if !isAlphaNumeric(gameName) {
			foundError++
			errorString += "Game name should be alphanumeric."
		}

		if len(playerName) > PLAYER_NAME_LENGTH || len(playerName) <= 0 {
			foundError++
			errorString += "Invalid game name length."
		}

		if !isAlpha(playerName) {
			foundError++
			errorString += "Game name should have only alphabets."
		}

		if foundError > 0 {
			resp := make(map[string]string)
			resp["status"] = "bad"
			resp["Errors"] = errorString
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(resp)
		} else {
			s.ServeHTTP(rw, r)
		}
	})
}
