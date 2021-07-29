package storewrapper

import (
	"github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"
)

type TtlDecoratorr struct {
	ttl        int //time in seconds
	wrappedObj keyvaluestore.KeyValueStore
}

// Sets a key value pair, if given ttl is 0, the decorator ttl is used
func (wrp *TtlDecoratorr) Set(key string, value string, ttl int) error {
	if ttl == 0 {
		ttl = int(wrp.ttl)
	}
	return wrp.wrappedObj.Set(key, value, ttl)
}

// Gets the value based on key, updated ttl as well
func (wrp *TtlDecoratorr) Get(key string) (string, error) {
	value, err := wrp.wrappedObj.Get(key)
	if err == nil {
		wrp.Set(key, value, wrp.ttl)
	}
	return value, err
}

// Update existing key value
func (wrp *TtlDecoratorr) Update(key string, value string, ttl int) (string, error) {
	return wrp.wrappedObj.Update(key, value, ttl)
}

// Delete key value store
func (wrp *TtlDecoratorr) Delete(key string) {
	wrp.wrappedObj.Delete(key)
}

// Connect to database/store
func (wrp *TtlDecoratorr) Connect() {
	wrp.wrappedObj.Connect()
}

// Disconnect from database
func (wrp *TtlDecoratorr) Disconnect() {
	wrp.wrappedObj.Disconnect()
}

func NewTtlDeocrator(store keyvaluestore.KeyValueStore, t int) *TtlDecoratorr {
	return &TtlDecoratorr{t, store}
}
