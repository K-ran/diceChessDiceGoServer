package gamehandler

import (
	"crypto/rand"
	"math/big"
)

// stores information about a single die
type diceStruct struct {
	Value int `json:"value"`
}

func (d *diceStruct) Rollit() {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(6)))
	d.Value = int(bigInt.Int64())
}

// stores information about the player
type player struct {
	Name string `json:"name"`
}

type GameState int

const (
	WAITING GameState = iota
	RUNNING
)

//stores state of the game
type diceChessGame struct {
	RespType  int          `json:"type"` //type of response
	Player1   player       `json:"p1"`
	Player2   player       `json:"p2"`
	GameId    string       `json:"gameID"`
	GameName  string       `json:"gameName"`
	State     GameState    `json:"gameState"`
	Dice      []diceStruct `json:"dice"`
	RollCount int          `json:"rollCount"`
}

func NewDiceChessGame(gameId string, gameName string, p1 string, p2 string, dice int) *diceChessGame {
	newGame := &diceChessGame{RESPONSE_GAME, player{p1}, player{p2}, gameId, gameName, WAITING, make([]diceStruct, 3), 0}
	return newGame
}

func (game *diceChessGame) SetPlayer1(player1 player) {
	game.Player1 = player1
}

func (game *diceChessGame) SetPlayer2(player2 player) {
	game.Player2 = player2
}

func (game *diceChessGame) SetGameId(id string) {
	game.GameId = id
}
