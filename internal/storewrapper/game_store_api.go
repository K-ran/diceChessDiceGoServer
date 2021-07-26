package storewrapper

import (
	"crypto/rand"
	"sync"

	"math/big"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

var player_mu sync.Mutex

const GAME_ID_LENGTH = 7
const PLAYER_ID_LENGTH = 32

//Generates a unique player id in store
func GreateUniquePlayerId(store keyvaluestore.KeyValueStore) string {
	value, _ := generateRandomString(PLAYER_ID_LENGTH)
	player_mu.Lock()
	defer player_mu.Unlock()

	_, err := store.Get(value)

	//keep trying till we get a unique key
	for err == nil {
		value, _ = generateRandomString(PLAYER_ID_LENGTH)
		_, err = store.Get(value)
	}

	store.Set(value, "", 0)

	return value
}

var game_mu sync.Mutex

//Generates a unique game id in store
//TODO: add another prefix character based in env variable GO_SEVER_PREFIX
func GreateUniqueGameId(store keyvaluestore.KeyValueStore) string {
	value, _ := generateRandomString(GAME_ID_LENGTH)
	game_mu.Lock()
	defer game_mu.Unlock()

	_, err := store.Get(value)

	//keep trying till we get a unique key
	for err == nil {
		value, _ = generateRandomString(GAME_ID_LENGTH)
		_, err = store.Get(value)
	}

	store.Set(value, "{}", 0)

	return value
}
