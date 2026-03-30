
type DLinkedList struct {
	key  int
	val  int
	prev *DLinkedList
	next *DLinkedList
}

type LRUCache struct {
	capacity int
	head     *DLinkedList
	tail     *DLinkedList
	cache    map[int]*DLinkedList
}

func Constructor(capacity int) LRUCache {
	// declare dummy nodes
	head, tail := &DLinkedList{}, &DLinkedList{}
	head.next = tail // head → tail
	tail.prev = head // head ← tail
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*DLinkedList),
	}
}

func (this *LRUCache) Get(key int) int {

	if node, exist := this.cache[key]; exist {
		this.removeNode(node)
		this.addToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, exist := this.cache[key]; exist {
		// update new value
		node.val = value
		// update node to head
		this.removeNode(node)
		this.addToHead(node)
		return
	}
	newNode := &DLinkedList{key: key, val: value}
	this.cache[key] = newNode
	this.addToHead(newNode)

	if len(this.cache) > this.capacity {
		lruCache := this.removeTail()
		delete(this.cache, lruCache.key)
	}

}
func (this *LRUCache) removeTail() *DLinkedList {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
func (this *LRUCache) removeNode(node *DLinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) addToHead(node *DLinkedList) {

	node.next = this.head.next // node -> recentNode
	this.head.next.prev = node // recentNode <- node

	node.prev = this.head // node <- head (dummy node)
	this.head.next = node // head -> node
}