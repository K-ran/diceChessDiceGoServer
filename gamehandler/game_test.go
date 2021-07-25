package gamehandler

import (
	"encoding/json"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewDiceChessGame("abcd", "Ram", "Shyam", 3)
	if value, err := json.Marshal(game); err == nil {
		compareString := "{\"p1\":{\"name\":\"Ram\"},\"p2\":{\"name\":\"Shyam\"},\"gameID\":\"abcd\",\"dice\":[{\"value\":0},{\"value\":0},{\"value\":0}]}"
		if compareString != string(value) {
			t.Fatal("Comparision failed, expected {}\n got {}", compareString, string(value))
		}
	} else {
		t.Fatal("Failed to marsh game object")
	}
}
