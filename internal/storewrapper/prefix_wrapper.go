package storewrapper

import "github.com/k-ran/diceChessDiceServer/pkg/keyvaluestore"

//Adds prefix to keys before sending to to key value store
type PrefixDecorator struct {
	wrappedObj keyvaluestore.KeyValueStore
	prefix     string
}

// Sets a key value pair
func (wrp *PrefixDecorator) Set(key string, value string, ttl int) error {
	return wrp.wrappedObj.Set(wrp.prefix+key, value, ttl)
}

// Gets the value based on key
func (wrp *PrefixDecorator) Get(key string) (string, error) {
	return wrp.wrappedObj.Get(wrp.prefix + key)
}

// Update existing key value
func (wrp *PrefixDecorator) Update(key string, value string, ttl int) (string, error) {
	return wrp.wrappedObj.Update(wrp.prefix+key, value, ttl)
}

// Delete key value store
func (wrp *PrefixDecorator) Delete(key string) {
	wrp.wrappedObj.Delete(wrp.prefix + key)
}

// Connect to database/store
func (wrp *PrefixDecorator) Connect() {
	wrp.wrappedObj.Connect()
}

// Disconnect from database
func (wrp *PrefixDecorator) Disconnect() {
	wrp.wrappedObj.Disconnect()
}

func NewPrefixDecorator(store keyvaluestore.KeyValueStore, prefix string) *PrefixDecorator {
	instance := &PrefixDecorator{store, prefix}
	return instance
}
