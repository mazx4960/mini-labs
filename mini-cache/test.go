package main

// import (
// 	"testing"
// 	"minicache/lfu"
// )

// func TestCreateCache(t testing.T) {
// 	cache := lfu.New(2)
// 	if cache.capacity != 2 {
// 		t.Errorf("expected capacity to be 2, but got %d", cache.capacity)
// 	}
// 	if cache.freqHead.freq != 0 {
// 		t.Errorf("expected freqHead.freq to be 0, but got %d", cache.freqHead.freq)
// 	}
// 	if len(cache.keyMap) != 0 {
// 		t.Errorf("expected len(keyMap) to be 0, but got %d", len(cache.keyMap))
// 	}
// }