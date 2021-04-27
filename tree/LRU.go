package tree

import "fmt"

type DListNode struct {
    key, value int
    next *DListNode
    pre *DListNode
}

type LRUCache struct {
	 capacity  int
	 cache map[int]*DListNode
	 head, tail *DListNode
}

func InitDListNode(key, value int)  *DListNode{
	return &DListNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru :=  LRUCache{
		capacity: capacity,
		cache: make(map[int]*DListNode, 0),
		head: InitDListNode(-1, -1),
		tail: InitDListNode(-1, -1 ),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}


func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		this.MoveToHead(value)
		return value.value
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		node := InitDListNode(key, value)
		if len(this.cache) < this.capacity {
			this.InsertToHead(node)
		} else {
			this.RemoveTail()
			this.InsertToHead(node)
		}
		this.cache[key] = node
	}
}

func (this *LRUCache) InsertToHead(node *DListNode) {
	p := this.head.next
	this.head.next = node
	p.pre = node
	node.pre = this.head
	node.next = p
}

func (this *LRUCache) MoveToHead(node *DListNode) {
	// delete node
	p1 := node.pre
	p2 := node.next
	p1.next = p2
	p2.pre = p1

	// insert to head
	this.InsertToHead(node)
}

func (this *LRUCache) RemoveTail() {
	// get need delete node
	node :=  this.tail.pre
	p1 := node.pre

	p1.next = this.tail
	this.tail.pre = p1

	delete(this.cache, node.key)
}

func (this *LRUCache)  PrintLRU() {
	node := this.head.next
	for node != nil {
		fmt.Printf("%d ->", node.value)
		node = node.next
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
