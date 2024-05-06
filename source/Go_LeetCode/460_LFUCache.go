package Go_LeetCode

//
//type LFUCache struct {
//	capacity int                 // 缓存的大小限制，即K
//	size     int                 // 缓存目前有多少个节点
//	records  map[int]*Node       // 表示key(Integer)由哪个节点(Node)代表
//	heads    map[*Node]*NodeList // 表示节点(Node)在哪个桶(NodeList)里
//	headList *NodeList           // 整个结构中位于最左的桶
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		capacity: capacity,
//		size:     0,
//		records:  make(map[int]*Node),
//		heads:    make(map[*Node]*NodeList),
//		headList: nil,
//	}
//}
//
//// 节点的数据结构
//type Node struct {
//	key   int
//	value int
//	times int   // 这个节点发生get或者set的次数总和
//	up    *Node // 节点之间是双向链表所以有上一个节点
//	down  *Node // 节点之间是双向链表所以有下一个节点
//}
//
//func NewNode(k, v, t int) *Node {
//	return &Node{
//		key:   k,
//		value: v,
//		times: t,
//	}
//}
//
//// 桶结构
//type NodeList struct {
//	head *Node     // 桶的头节点
//	tail *Node     // 桶的尾节点
//	last *NodeList // 桶之间是双向链表所以有前一个桶
//	next *NodeList // 桶之间是双向链表所以有后一个桶
//}
//
//func newNodeList(node *Node) *NodeList {
//	return &NodeList{
//		head: node,
//		tail: node,
//	}
//}
//
//// AddNodeFromHead 把一个新的节点加入这个桶，新的节点都放在顶端变成新的头部
//func (list *NodeList) AddNodeFromHead(newHead *Node) {
//	newHead.down = list.head
//	list.head.up = newHead
//	list.head = newHead
//}
//
//// IsEmpty 判断这个桶是不是空的
//func (list *NodeList) IsEmpty() bool {
//	return list.head == nil
//}
//
//// DeleteNode 删除node节点并保证node的上下环境重新连接
//func (list *NodeList) DeleteNode(node *Node) {
//	if list.head == list.tail {
//		list.head = nil
//		list.tail = nil
//	} else {
//		if node == list.head {
//			list.head = node.down
//			list.head.up = nil
//		} else if node == list.tail {
//			list.tail = node.up
//			list.tail.down = nil
//		} else {
//			node.up.down = node.down
//			node.down.up = node.up
//		}
//	}
//	node.up = nil
//	node.down = nil
//}
//
//// ModifyHeadList
//// removeNodeList：刚刚减少了一个节点的桶
//// 这个函数的功能是，判断刚刚减少了一个节点的桶是不是已经空了。
//// 1）如果不空，什么也不做
////
//// 2)如果空了，removeNodeList还是整个缓存结构最左的桶(headList)。
//// 删掉这个桶的同时也要让最左的桶变成removeNodeList的下一个。
////
//// 3)如果空了，removeNodeList不是整个缓存结构最左的桶(headList)。
//// 把这个桶删除，并保证上一个的桶和下一个桶之间还是双向链表的连接方式
////
//// 函数的返回值表示刚刚减少了一个节点的桶是不是已经空了，空了返回true；不空返回false
//func (cache *LFUCache) ModifyHeadList(removeNodeList *NodeList) bool {
//	if removeNodeList.IsEmpty() {
//		if cache.headList == removeNodeList {
//			cache.headList = removeNodeList.next
//			if cache.headList != nil {
//				cache.headList.last = nil
//			}
//		} else {
//			removeNodeList.last.next = removeNodeList.next
//			if removeNodeList.next != nil {
//				removeNodeList.next.last = removeNodeList.last
//			}
//		}
//		return true
//	}
//	return false
//}
//
//// Move 函数的功能
//// node这个节点的次数+1了，这个节点原来在oldNodeList里。
//// 把node从oldNodeList删掉，然后放到次数+1的桶中
//// 整个过程既要保证桶之间仍然是双向链表，也要保证节点之间仍然是双向链表
//func (cache *LFUCache) Move(node *Node, oldNodeList *NodeList) {
//	oldNodeList.DeleteNode(node)
//	// preList表示次数+1的桶的前一个桶是谁
//	// 如果oldNodeList删掉node之后还有节点，oldNodeList就是次数+1的桶的前一个桶
//	// 如果oldNodeList删掉node之后空了，oldNodeList是需要删除的，所以次数+1的桶的前一个桶，是oldNodeList的前一个
//	var preList *NodeList
//	if cache.ModifyHeadList(oldNodeList) {
//		preList = oldNodeList.last
//	} else {
//		preList = oldNodeList
//	}
//	// nextList表示次数+1的桶的后一个桶是谁
//	nextList := oldNodeList.next
//	if nextList == nil {
//		newList := newNodeList(node)
//		if preList != nil {
//			preList.next = newList
//		}
//		newList.last = preList
//		if cache.headList == nil {
//			cache.headList = newList
//		}
//		cache.heads[node] = newList
//	} else {
//		if nextList.head.times == node.times {
//			nextList.AddNodeFromHead(node)
//			cache.heads[node] = nextList
//		} else {
//			newList := newNodeList(node)
//			if preList != nil {
//				preList.next = newList
//			}
//			newList.last = preList
//			newList.next = nextList
//			nextList.last = newList
//			if cache.headList == nextList {
//				cache.headList = newList
//			}
//			cache.heads[node] = newList
//		}
//	}
//}
//
//func (this *LFUCache) Put(key int, value int) {
//	if this.capacity == 0 {
//		return
//	}
//	if node, ok := this.records[key]; ok {
//		node.value = value
//		node.times++
//		curNodeList := this.heads[node]
//		this.Move(node, curNodeList)
//	} else {
//		if this.size == this.capacity {
//			node := this.headList.tail
//			this.headList.DeleteNode(node)
//			this.ModifyHeadList(this.headList)
//			delete(this.records, node.key)
//			delete(this.heads, node)
//			this.size--
//		}
//		node := NewNode(key, value, 1)
//		if this.headList == nil {
//			this.headList = newNodeList(node)
//		} else {
//			if this.headList.head.times == node.times {
//				this.headList.AddNodeFromHead(node)
//			} else {
//				newList := newNodeList(node)
//				newList.next = this.headList
//				this.headList.last = newList
//				this.headList = newList
//			}
//		}
//		this.records[key] = node
//		this.heads[node] = this.headList
//		this.size++
//	}
//}
//
//func (this *LFUCache) Get(key int) int {
//	if node, ok := this.records[key]; !ok {
//		return -1
//	} else {
//		node.times++
//		curNodeList := this.heads[node]
//		this.Move(node, curNodeList)
//		return node.value
//	}
//}
