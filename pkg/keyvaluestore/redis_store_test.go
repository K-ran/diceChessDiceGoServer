package keyvaluestore

import "testing"

func TestRedisStore(t *testing.T) {
	rdb := NewRedisStore()
	if rdb.Set("name", "shyam", 20) != nil {
		t.Fatal("Set Fail")
	}

	if value, _ := rdb.Get("name"); value != "shyam" {
		t.Fatal("Get Failed")
	}

	if _, err := rdb.Update("name", "ram", 20); err != nil {
		t.Fatal("Update failed")
	}

	if value, _ := rdb.Get("name"); value != "ram" {
		t.Fatal("Update failed")
	}
}
