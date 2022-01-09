package lfu

type LFUCache struct {
	capacity int // max number of items in the cache
	freqHead *freqNode
	keyMap map[string]*cacheItem // key -> cacheItem
}

type freqNode struct {
	freq int
	items map[string]bool // workaround for a set of keys
	prev *freqNode
	next *freqNode
}

type cacheItem struct {
	key string
	val string
	parent *freqNode
}