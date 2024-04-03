package Go_LeetCode

import "fmt"

/*
已知一个消息流会不断地吐出整数1~N，但不一定按照顺序依次吐出，
如果上次打印的序号为i， 那么当i+1出现时
请打印i+1及其之后接收过的并且连续的所有数，直到1~N全部接收并打印完，
请设计这种接收并打印的结构
*/

type Node_02_03 struct {
	info string
	next *Node_02_03
}

type MessageBox struct {
	headMap   map[int]*Node_02_03
	tailMap   map[int]*Node_02_03
	waitPoint int
}

func NewMessageBox() *MessageBox {
	return &MessageBox{
		headMap:   make(map[int]*Node_02_03),
		tailMap:   make(map[int]*Node_02_03),
		waitPoint: 1,
	}
}

func (mb *MessageBox) Receive(num int, info string) {
	if num < 1 {
		return
	}
	cur := &Node_02_03{info: info}
	mb.headMap[num] = cur
	mb.tailMap[num] = cur

	if preTail, ok := mb.tailMap[num-1]; ok {
		preTail.next = cur
		delete(mb.tailMap, num-1)
		delete(mb.headMap, num)
	}

	if nextHead, ok := mb.headMap[num+1]; ok {
		cur.next = nextHead
		delete(mb.headMap, num+1)
		delete(mb.tailMap, num)
	}

	if num == mb.waitPoint {
		mb.Print()
	}
}

func (mb *MessageBox) Print() {
	node := mb.headMap[mb.waitPoint]
	delete(mb.headMap, mb.waitPoint)

	for node != nil {
		fmt.Print(node.info, " ")
		node = node.next
		mb.waitPoint++
	}
	delete(mb.tailMap, mb.waitPoint-1)
	fmt.Println()
}
