package KeyValueStore

type KeyError struct {
	ErrString string
}

func (err *KeyError) Error() string {
	return err.ErrString
}

// Interface for a Game Storage unit.
type KeyValueStore interface {

	// Sets a key value pair
	Set(key string, value string)

	// Gets the value based on key
	Get(key string) (string, error)

	// Update existing key value
	Update(key string, value string) (string, error)

	// Delete key value store
	Delete(key string)

	// Connect to database/store
	Connect() error

	// Disconnect from database
	Disconnect() error
}
