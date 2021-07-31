package storewrapper

import (
	"crypto/rand"
	"os"
	"strconv"
	"strings"
	"sync"

	"math/big"

	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

const GAME_ID_LENGTH = 8
const PLAYER_ID_LENGTH = 32
const LETTERS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

// Generates base64 random string of length n
func generateRandomString(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(LETTERS))))
		if err != nil {
			return "", err
		}
		ret[i] = LETTERS[num.Int64()]
	}

	return string(ret), nil
}

var player_mu sync.Mutex

//Generates a unique player id in store
func GreateUniquePlayerId(store keyvaluestore.KeyValueStore, gameId string) string {
	value, _ := generateRandomString(PLAYER_ID_LENGTH)
	player_mu.Lock()
	defer player_mu.Unlock()

	_, err := store.Get(value)

	//keep trying till we get a unique key
	for err == nil {
		value, _ = generateRandomString(PLAYER_ID_LENGTH)
		_, err = store.Get(value)
	}

	store.Set(value, gameId, 0)

	return value
}

var game_mu sync.Mutex

// Generated game prefix based on Env variable
func getRandomServerPrefix() string {
	prefix := os.Getenv("GO_SERVER_PEFIX")
	if prefix == "" {
		prefix = "0-63"
	}

	start, _ := strconv.Atoi(strings.Split(prefix, "-")[0])
	end, _ := strconv.Atoi(strings.Split(prefix, "-")[1])
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(end-start+1)))
	prefixString := string(LETTERS[start+int(num.Int64())])
	return prefixString
}

//Generates a unique game id in store
//TODO: add another prefix character based in env variable GO_SEVER_PREFIX
func GreateUniqueGameId(store keyvaluestore.KeyValueStore) string {
	value, _ := generateRandomString(GAME_ID_LENGTH - 1)
	game_mu.Lock()
	defer game_mu.Unlock()

	_, err := store.Get(value)

	//keep trying till we get a unique key
	for err == nil {
		value, _ = generateRandomString(GAME_ID_LENGTH - 1)
		_, err = store.Get(value)
	}
	prefix := getRandomServerPrefix()
	value = prefix + value
	store.Set(value, "{}", 0)

	return value
}
