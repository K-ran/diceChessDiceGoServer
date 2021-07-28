package gamehandler

import (
	"encoding/json"
	"net/http"
)

//Create a failure response status
func ReturnFailure(rw http.ResponseWriter, err string) {
	resp := make(map[string]string)
	resp["status"] = "1"
	resp["Errors"] = err
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(resp)
}
