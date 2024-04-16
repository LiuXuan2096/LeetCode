package Go_LeetCode

// 这道题没做完，没看懂讲解 先不做了。
import "sort"

type Node_218 struct {
	x     int  // 起点或终点的坐标位置
	isAdd bool // 是加高度还是减高度
	h     int  // 该点对应的高度
}

type NodeSlice []Node_218

func (ns NodeSlice) Len() int {
	return len(ns)
}

func (ns NodeSlice) Less(i, j int) bool {
	return ns[i].x < ns[j].x
}

func (ns NodeSlice) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func getSkyline(buildings [][]int) [][]int {
	// 大楼数量*2就是对象数，因为每个大楼都有起点和终点以及对应的高度
	nodes := make([]Node_218, len(buildings)*2)
	for i := 0; i < len(buildings); i++ {
		nodes[i*2] = Node_218{buildings[i][0], true, buildings[i][2]}
		nodes[i*2+1] = Node_218{buildings[i][1], false, buildings[i][2]}

	}
	sort.Sort(NodeSlice(nodes))

	mapHeightTimes := make(map[int]int) //key是高度 value是出现的次数
	mapXHeight := make(map[int]int)     //每一个x的最大高度是多少

	for i := 0; i < len(nodes); i++ {
		if nodes[i].isAdd {
			mapHeightTimes[nodes[i].h]++
		} else {
			if mapHeightTimes[nodes[i].h] == 1 {
				// 某个高度为出现次数为0时要从map中删掉 否则会干扰排序
				delete(mapHeightTimes, nodes[i].h)
			} else {
				mapHeightTimes[nodes[i].h]--
			}
		}

		if len(mapHeightTimes) == 0 {
			mapXHeight[nodes[i].x] = 0
		} else {
			maxHeight := 0
			for h := range mapHeightTimes {
				maxHeight = max_218(maxHeight, h)
			}
			mapXHeight[nodes[i].x] = maxHeight
		}
	}

	ans := [][]int{}
	for x, maxHeight := range mapXHeight {
		if len(ans) == 0 || ans[len(ans)-1][1] != maxHeight {
			ans = append(ans, []int{x, maxHeight})
		}
	}
	return ans
}

func max_218(a, b int) int {
	if a > b {
		return a
	}
	return b
}
