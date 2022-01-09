package main

import (
	"log"
	"minicache/lfu"
)

func main() {
	cache := lfu.New(2)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	log.Println(cache.Get("key1"))

	// key2 should be evicted
	cache.Put("key3", "value3")

	log.Println(cache.Get("key1"))
	log.Println(cache.Get("key2"))
	log.Println(cache.Get("key3"))
}