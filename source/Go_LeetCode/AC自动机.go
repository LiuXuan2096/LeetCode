package Go_LeetCode

import "container/list"

// ACNode
//前缀树的节点，如果一个ACNode，end为空，不是结尾。
//如果end不为空，表示这个节点是某个字符串的结尾，end的值就是这个字符串
//只有在end变量不为空的时候，endUse才有意义：这个字符串之前没有加入过答案
///*
type ACNode struct {
	end    string
	endUse bool
	fail   *ACNode
	nexts  []*ACNode
}

type ACAutomation struct {
	root *ACNode
}

func NewACAutomation() *ACAutomation {
	return &ACAutomation{root: &ACNode{nexts: make([]*ACNode, 26)}}
}

func (ac *ACAutomation) Insert(s string) {
	str := []rune(s)
	cur := ac.root
	var index int
	for i := 0; i < len(str); i++ {
		index = int(str[i] - 'a')
		if cur.nexts[index] == nil {
			cur.nexts[index] = &ACNode{nexts: make([]*ACNode, 26)}
		}
		cur = cur.nexts[index]
	}
	cur.end = s
}

func (ac *ACAutomation) Build() {
	queue := list.New()
	queue.PushBack(ac.root)
	var cur, cfail *ACNode
	for queue.Len() > 0 {
		cur = queue.Remove(queue.Front()).(*ACNode)
		for i := 0; i < 26; i++ {
			if cur.nexts[i] != nil {
				cur.nexts[i].fail = ac.root
				cfail = cur.fail
				for cfail != nil {
					if cfail.nexts[i] != nil {
						cur.nexts[i].fail = cfail.nexts[i]
						break
					}
					cfail = cfail.fail
				}
				queue.PushBack(cur.nexts[i])
			}
		}
	}
}

func (ac *ACAutomation) ContainWords(content string) []string {
	str := []rune(content)
	cur := ac.root
	var follow *ACNode
	var index int
	var ans []string

	for i := 0; i < len(str); i++ {
		index = int(str[i] - 'a')
		// 如果当前字符在这条路上没配出来，就随着fail方向走向下条路径
		for cur.nexts[index] == nil && cur != ac.root {
			cur = cur.fail
		}
		// 区分两种情况：
		// 1) 现在来到的路径，是可以继续匹配的
		// 2) 现在来到的节点，就是前缀树的根节点
		cur = func() *ACNode {
			if cur.nexts[index] != nil {
				return cur.nexts[index]
			}
			return ac.root
		}()
		follow = cur
		for follow != ac.root {
			if follow.endUse {
				break
			}
			if follow.end != "" {
				// 不同的需求，在这一段之间修改
				ans = append(ans, follow.end)
				follow.endUse = true
			}
			follow = follow.fail
		}
	}

	return ans
}
