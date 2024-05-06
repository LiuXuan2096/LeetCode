package Go_LeetCode

//
//type LRUCache struct {
//	cache *MyCache
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		cache: NewMyCache(capacity),
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	ans := this.cache.Get(key)
//	if ans == nil {
//		return -1
//	}
//	return ans.(int)
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	this.cache.Set(key, value)
//}
//
///**
// * Your LRUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
//
//type Node struct {
//	key   int
//	value int
//	last  *Node //双链表上一个节点
//	next  *Node //双链表下一个节点
//}
//
//func NewNode(key, value int) *Node {
//	return &Node{
//		key:   key,
//		value: value,
//	}
//}
//
//type NodeDoubleLinkedList struct {
//	head *Node
//	tail *Node
//}
//
//func NewNodeDoubleLinkedList() *NodeDoubleLinkedList {
//	return &NodeDoubleLinkedList{}
//}
//
//// AddNode 现在来了一个新的node，请挂到尾巴上去
//// 双链表的尾结点放最新的node 头结点放最旧的node
//func (list *NodeDoubleLinkedList) AddNode(newNode *Node) {
//	if newNode == nil {
//		return
//	}
//	if list.head == nil {
//		list.head = newNode
//		list.tail = newNode
//	} else {
//		list.tail.next = newNode
//		newNode.last = list.tail
//		list.tail = newNode
//	}
//}
//
//// MoveNodeToTail 更新node时，将更新后的node从双链表的原始位置挪到双链表的尾部
//// node 入参，一定保证！node在双向链表里！
//// node原始的位置，左右重新连好，然后把node分离出来
//// 挂到整个链表的尾巴上
//func (list *NodeDoubleLinkedList) MoveNodeToTail(node *Node) {
//	if list.tail == node {
//		return
//	}
//	if list.head == node {
//		list.head = node.next
//		list.head.last = nil
//	} else {
//		node.last.next = node.next
//		node.next.last = node.last
//	}
//	node.last = list.tail
//	node.next = nil
//	list.tail.next = node
//	list.tail = node
//}
//
//func (list *NodeDoubleLinkedList) RemoveHead() *Node {
//	if list.head == nil {
//		return nil
//	}
//	res := list.head
//	if list.head == list.tail {
//		list.head = nil
//		list.tail = nil
//	} else {
//		list.head = res.next
//		res.next = nil
//		list.head.last = nil
//	}
//	return res
//}
//
//type MyCache struct {
//	keyNodeMap map[int]*Node
//	nodeList   *NodeDoubleLinkedList
//	capacity   int
//}
//
//func NewMyCache(capacity int) *MyCache {
//	return &MyCache{
//		keyNodeMap: make(map[int]*Node),
//		nodeList:   NewNodeDoubleLinkedList(),
//		capacity:   capacity,
//	}
//}
//
//func (cache *MyCache) Get(key int) interface{} {
//	if node, ok := cache.keyNodeMap[key]; ok {
//		cache.nodeList.MoveNodeToTail(node)
//		return node.value
//	}
//	return nil
//}
//
//// Set(Key, Value)
//// 新增  更新value的操作
//func (cache *MyCache) Set(key, value int) {
//	if node, ok := cache.keyNodeMap[key]; ok {
//		node.value = value
//		cache.nodeList.MoveNodeToTail(node)
//	} else {
//		newNode := NewNode(key, value)
//		cache.keyNodeMap[key] = newNode
//		cache.nodeList.AddNode(newNode)
//		if len(cache.keyNodeMap) == cache.capacity+1 {
//			cache.RemoveMostUnUsedCache()
//		}
//	}
//}
//
//func (cache *MyCache) RemoveMostUnUsedCache() {
//	removeNode := cache.nodeList.RemoveHead()
//	delete(cache.keyNodeMap, removeNode.key)
//}
