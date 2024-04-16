package Go_LeetCode

import "math"

// 前缀树的Node结构
// nexts[0] -> 0方向的路
// nexts[1] -> 1方向的路
// nexts[0] == null 0方向上没路！
// nexts[0] != null 0方向有路，可以跳下一个节点
// nexts[1] == null 1方向上没路！
// nexts[1] != null 1方向有路，可以跳下一个节点
type Node_421 struct {
	nexts [2]*Node_421
}

type NumTrie struct {
	head *Node_421
}

func NewNumTrie() *NumTrie {
	return &NumTrie{head: &Node_421{}}
}

func (t *NumTrie) add(newNum int) {
	cur := t.head
	for move := 31; move >= 0; move-- {
		path := (newNum >> move) & 1
		if cur.nexts[path] == nil {
			cur.nexts[path] = &Node_421{}
		}
		cur = cur.nexts[path]
	}
}

// maxXor 该结构之前收集了一票数字，并且建好了前缀树
// num和 谁 ^ 最大的结果（把结果返回）
func (t *NumTrie) maxXor(num int) int {
	cur := t.head
	ans := 0
	for move := 31; move >= 0; move-- {
		// 取出num中第move位的状态，path只有两种值0就1，整数
		path := (num >> move) & 1
		// 期待遇到的东西
		var best int
		if move == 31 {
			best = path //符号位想遇到和自己一样的
		} else {
			best = path ^ 1 //非符号位想遇到和自己不一样的
		}
		// 实际遇到的东西
		if cur.nexts[best] != nil {
			best = best
		} else {
			best = best ^ 1
		}
		// (path ^ best) 当前位位异或完的结果
		ans |= (path ^ best) << move
		cur = cur.nexts[best]
	}
	return ans
}

func MaxXorSubarray2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := math.MinInt32
	xor := 0 // 存放0~i整体异或和
	numTire := NewNumTrie()
	numTire.add(0)

	for i := 0; i < len(arr); i++ {
		xor ^= arr[i]
		if xorMax := numTire.maxXor(xor); xorMax > max {
			max = xorMax
		}
	}
	return max
}

func findMaximumXOR(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	max := math.MinInt32
	numTrie := NewNumTrie()
	numTrie.add(nums[0])
	for i := 1; i < len(nums); i++ {
		if xorMax := numTrie.maxXor(nums[i]); xorMax > max {
			max = xorMax
		}
		numTrie.add(nums[i])
	}
	return max
}
