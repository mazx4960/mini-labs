package lfu

import "errors"

func New(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		freqHead: NewFreqNode(),
		keyMap: make(map[string]*cacheItem),
	}
}

func (cache *LFUCache) Get(key string) (string, error) {
	if node, ok := cache.keyMap[key]; ok {
		cache.incrementFreq(node)
		return node.val, nil
	}
	return "", errors.New("key not found")
}

func (cache *LFUCache) Put(key string, val string) error {
	if cache.capacity == 0 {
		return errors.New("cache capacity is 0")
	}

	if node, ok := cache.keyMap[key]; ok {
		node.val = val
		cache.incrementFreq(node)
		return nil
	}

	cache.Insert(key, val)
	return nil
}

func (cache *LFUCache) Insert(key string, val string) error {
	if cache.capacity == 0 {
		return errors.New("cache capacity is 0")
	}

	if _, ok := cache.keyMap[key]; ok {
		return errors.New("key already exists")
	}

	if len(cache.keyMap) == cache.capacity {
		cache.evict()
	}

	node := NewCacheItem(key, val, cache.freqHead)
	cache.keyMap[key] = node
	cache.incrementFreq(node)
	return nil
}

func (cache *LFUCache) Delete(key string) error {
	if node, ok := cache.keyMap[key]; ok {
		DeleteCacheItemFromFreq(node)
		delete(cache.keyMap, node.key)
		return nil
	}
	return errors.New("key not found")
}

func (cache *LFUCache) incrementFreq(node *cacheItem) {
	freq := node.parent
	nextFreq := freq.next

	if nextFreq == cache.freqHead || nextFreq.freq != freq.freq + 1 {
		nextFreq = GetNewFreqNode(freq.freq + 1, freq, nextFreq)
	}
	DeleteCacheItemFromFreq(node)
	node.parent = nextFreq
	nextFreq.items[node.key] = true
}

// Evict removes the least frequently used item from the cache.
func (cache *LFUCache) evict() error {
	leastFreq := cache.freqHead.next
	if leastFreq == cache.freqHead {
		return errors.New("cache is empty")
	}

	key := GetFirstKey(leastFreq.items)
	cache.Delete(key)
	return nil
}
