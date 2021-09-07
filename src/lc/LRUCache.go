package lc

type LRUCache struct {
	head      *ListNodeLRU
	tail      *ListNodeLRU
	numToNode map[int]*ListNodeLRU
	capacity  int
}

type ListNodeLRU struct {
	key   int
	value int
	next  *ListNodeLRU
	prev  *ListNodeLRU
}

func constructListNodeLRU(key, value int) *ListNodeLRU {
	return &ListNodeLRU{key: key, value: value}
}

func ConstructorLRU(capacity int) LRUCache {
	location := make(map[int]*ListNodeLRU)
	head := &ListNodeLRU{-1, -1, nil, nil}
	tail := &ListNodeLRU{-1, -1, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{numToNode: location, head: head, tail: tail, capacity: capacity}

}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.numToNode[key]; !ok {
		return -1
	}
	this.moveToTail(key)

	return this.numToNode[key].value

}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.numToNode[key]; ok {
		this.numToNode[key].value = value
		this.numToNode[key].key = key
		this.moveToTail(key)
	}
	if len(this.numToNode) < this.capacity {
		this.addToTail(key, value)
	} else {
		this.deleteHead(key, value)
	}

}

func (this *LRUCache) moveToTail(key int) {
	node := this.numToNode[key]
	if node == this.tail {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	this.tail.next = node
	node.prev = this.tail
	this.tail = node

}
func (this *LRUCache) addToTail(key, value int) {
	node := &ListNodeLRU{key: key, value: value}
	this.tail.next = node
	node.prev = this.tail
	this.tail = node
	this.numToNode[key] = node
}
func (this *LRUCache) deleteHead(key, value int) {
	delete(this.numToNode, this.head.next.key)
	this.head.next.key = key
	this.head.next.value = value
	this.numToNode[key] = this.head.next
	this.moveToTail(key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
