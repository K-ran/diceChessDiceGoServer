package gamehandler

import (
	"encoding/json"
	"net/http"
)

const (
	RESPONSE_ERROR = iota
	RESPONSE_GAME
	RESPONSE_CREATE
	RESPONSE_JOIN
	RESPONSE_GETSTATUS
	RESPONSE_ROLL
)

//Create a failure response status
func ReturnFailure(rw http.ResponseWriter, err string) {
	resp := make(map[string]interface{})
	resp["type"] = RESPONSE_ERROR
	resp["Errors"] = err
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(resp)
}
