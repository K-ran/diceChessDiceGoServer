package main

// stores information about a single die
type diceStruct struct {
	Value int `json:"value"`
}

// stores information about the player
type player struct {
	Name string `json:"name"`
}

//stores state of the game
type diceChessGame struct {
	Player1 player       `json:"p1"`
	Player2 player       `json:"p2"`
	GameId  string       `json:"gameID"`
	Dice    []diceStruct `json:"dice"`
}

func NewDiceChessGame(gameId string, p1 string, p2 string, dice int) *diceChessGame {
	newGame := &diceChessGame{player{p1}, player{p2}, gameId, make([]diceStruct, 3)}
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
