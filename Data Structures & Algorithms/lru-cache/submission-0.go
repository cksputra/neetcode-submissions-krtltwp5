type ListNode struct {
    Key int
    Val int
    Next *ListNode
    Prev *ListNode
}


type LRUCache struct {
    Max int 
    Current int
    HashMap map[int]*ListNode
    Left *ListNode
    Right *ListNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        Max: capacity,
        HashMap: make(map[int]*ListNode),
        Left: &ListNode{},
        Right: &ListNode{},
    }
    lru.Left.Next = lru.Right
    lru.Right.Prev = lru.Left

    return lru
    
}

func (this *LRUCache) remove(node *ListNode) {
    prev, next := node.Prev, node.Next
    prev.Next = next
    next.Prev = prev
}

func (this *LRUCache) insert(node *ListNode) {
    prev, next := this.Right.Prev, this.Right
    prev.Next = node
    next.Prev = node
    node.Next = next
    node.Prev = prev
}


func (this *LRUCache) Get(key int) int {
    if v, ok := this.HashMap[key]; ok {
        this.remove(v)
        this.insert(v)
        return v.Val
    }

    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if v, ok := this.HashMap[key]; ok {
        this.remove(v)
        delete(this.HashMap, key)
    }

    node := &ListNode{
        Val: value,
        Key: key,
    }

    this.HashMap[key] = node
    this.insert(node)

    if len(this.HashMap) > this.Max {
        lru := this.Left.Next
        this.remove(lru)
        delete(this.HashMap, lru.Key)
    }
}
