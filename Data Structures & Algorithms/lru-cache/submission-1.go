type ListNode struct {
	Key int
 	Val int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
    Capacity int

	Cache map[int]*ListNode

	Left *ListNode
	Right *ListNode
}

func (this *LRUCache) insert(node *ListNode) {
	prev, next := this.Right.Prev, this.Right

	prev.Next = node
	next.Prev = node

	node.Next = next
	node.Prev = prev
}

func (this *LRUCache) remove(node *ListNode) {
	prev, next := node.Prev, node.Next

	prev.Next = next
	next.Prev = prev
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*ListNode)

	lru := LRUCache{
		Capacity: capacity,
		Cache: cache,
		Left: &ListNode{},
		Right: &ListNode{},
	}

	lru.Left.Next = lru.Right
	lru.Right.Prev = lru.Left
	return lru
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

	node := &ListNode{
		Val: value,
		Key: key,
	}

	this.Cache[key] = node
	this.insert(node)

	// eviction
	if len(this.Cache) > this.Capacity {
		lru := this.Left.Next
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}
}
