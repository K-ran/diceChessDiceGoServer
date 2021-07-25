package KeyValueStore

import "testing"

func TestRedisStore(t *testing.T) {
	rdb := NewRedisStore()
	if rdb.Set("name", "shyam") != nil {
		t.Fatal("Set Fail")
	}

	if value, _ := rdb.Get("name"); value != "shyam" {
		t.Fatal("Get Failed")
	}

	if _, err := rdb.Update("name", "ram"); err != nil {
		t.Fatal("Update failed")
	}

	if value, _ := rdb.Get("name"); value != "ram" {
		t.Fatal("Update failed")
	}
}
