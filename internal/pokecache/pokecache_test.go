package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second)
	if cache.cache == nil {
		t.Error("cache is empty")
	}

}
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Second)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, tcase := range cases {
		cache.Add(tcase.inputKey, tcase.inputVal)
		actual, ok := cache.Get(tcase.inputKey)
		if !ok {
			t.Errorf("%s not found", tcase.inputKey)
		}
		if string(actual) != string(tcase.inputVal) {
			t.Errorf("%s value doesn't match %s", string(actual), tcase.inputVal)
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("%s found value.", keyOne)
	}

}
func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(time.Millisecond)
	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("%s not found.", keyOne)
	}

}
