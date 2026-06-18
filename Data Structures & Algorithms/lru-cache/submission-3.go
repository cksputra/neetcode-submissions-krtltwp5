type Node struct {
	Key int
	Val int

	Next *Node
	Prev *Node
}

type LRUCache struct {
	capacity int

	Cache map[int]*Node

	Left *Node //pointer to lru
	Right *Node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		capacity: capacity,
		Cache: make(map[int]*Node),
		Left: &Node{},
		Right: &Node{},
	}

	lru.Left.Next = lru.Right
	lru.Right.Prev = lru.Left
	return lru
}



func (this *LRUCache) insert(node *Node) {
	prev, next := this.Right.Prev, this.Right

	node.Next = next
	node.Prev = prev

	prev.Next = node
	next.Prev = node
}

func (this *LRUCache) remove(node *Node){
	prev, next := node.Prev, node.Next
	
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) Get(key int) int {
    if v, ok := this.Cache[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if v, ok := this.Cache[key]; ok {
		this.remove(v)
		delete(this.Cache, key)
	}

	node := &Node{
		Val: value,
		Key: key,
	}

	this.Cache[key] = node
	this.insert(node)

	// eviction

	if len(this.Cache) > this.capacity {
		lru:= this.Left.Next
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}
}
