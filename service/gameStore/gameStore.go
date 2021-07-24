package main

type KeyError struct {
	ErrString string
}

func (err *KeyError) Error() string {
	return err.ErrString
}

// Interface for a Game Storage unit.
type GameStore interface {

	// Sets a key value pair
	SetGame(key string, value string)

	// Gets the value based on key
	GetGame(key string) (string, error)

	// Update existing key value
	UpdateGame(key string, value string) (string, error)

	// Delete key value store
	DeleteGame(key string)

	// Connect to database/store
	Connect() error
}
