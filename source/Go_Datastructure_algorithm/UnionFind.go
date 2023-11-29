package Go_LeetCode

const MAXN = 1000001

var father [MAXN]int
var size [MAXN]int
var help [MAXN]int // 在寻找集合代表点时把父子关系路径打平，让每次查找时 查找路径上的点的父亲节点都变成代表节点
var sets = 0       // 表示并查集的数量
// 初始化并查集
func InitUnionFind(n int) {
	for i := 0; i < n; i++ {
		father[i] = i
		size[i] = 1
	}
	sets = n
}

// 从i开始寻找集合代表点
func Find(i int) int {
	hi := 0
	// 一开始每个节点的父亲都是自己，如果一个i!=father[i]
	// 说明该节点的父亲不是自己
	for i != father[i] {
		help[hi] = i // 收集代表节点的子节点
		hi++
		i = father[i]
	}
	for hi--; hi >= 0; hi-- {
		father[help[hi]] = i // 把集合的代表节点 设为该集合节点的父亲
	}
	return i // 返回代表节点
}

// 查询x y是否是一个集合
func IsSameSet(x, y int) bool {
	return Find(x) == Find(y)
}

// x所在的集合 与y所在的集合 合并成一个集合
func Union(x, y int) {
	var fatherOfx = Find(x)
	var fatherOfy = Find(y)
	if fatherOfx != fatherOfy {
		if size[fatherOfx] >= size[fatherOfy] {
			size[fatherOfx] += size[fatherOfy]
			father[fatherOfy] = fatherOfx
		} else {
			size[fatherOfy] += size[fatherOfx]
			father[fatherOfx] = fatherOfy
		}
		sets--
	}

}

func Sets() int {
	return sets
}
