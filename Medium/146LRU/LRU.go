package _46LRU


type LRUCache struct {
	hash map[int]*DLinkedNode
	head, tail *DLinkedNode
	capacity int
	size int
}

// 双向链表，头结点为最新，尾节点为最久
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		hash: make(map[int]*DLinkedNode, 2),
		head: initDLinkedNode(0,0),
		tail: initDLinkedNode(0,0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// Get 取到key后需要将key放到链表头/**
func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	// 两种情况，map已存在和不存在，不存在需要判断当前cache大小，如果size==capacity则移除队尾，已存在则赋值并移动到队首
	if node, ok := this.hash[key]; ok {
		node.value = value
		this.moveToHead(node) //注意是移动，而不是添加
	} else {
		node = initDLinkedNode(key, value)
		this.hash[key] = node
		this.AddToHead(node)
		this.size++
		// 移除最久未使用
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.hash, removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) AddToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func remove(nums []int, val int) []int {
	length := len(nums)
	fast, slow := 0, 0
	for ;fast < length; fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return nums
}