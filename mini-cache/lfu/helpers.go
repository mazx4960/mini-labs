package lfu

func NewFreqNode() *freqNode {
	node := freqNode{
		freq: 0,
		items: make(map[string]bool),
		prev: nil,
		next: nil,
	}
	node.prev = &node
	node.next = &node
	return &node
}

func NewCacheItem(key string, val string, parent *freqNode) *cacheItem {
	return &cacheItem{
		key: key,
		val: val,
		parent: parent,
	}
}

func GetNewFreqNode(freq int, prev *freqNode, next *freqNode) *freqNode {
	node := freqNode{
		freq: freq,
		items: make(map[string]bool),
		prev: prev,
		next: next,
	}
	prev.next = &node
	next.prev = &node
	return &node
}

func DeleteFreqNode(node *freqNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func DeleteCacheItemFromFreq(node *cacheItem) {
	freq := node.parent
	delete(freq.items, node.key)
	if len(freq.items) == 0 {
		DeleteFreqNode(freq)
	}
}

func GetFirstKey(m map[string]bool) string {
	for k := range m {
		return k
	}
	return ""
}