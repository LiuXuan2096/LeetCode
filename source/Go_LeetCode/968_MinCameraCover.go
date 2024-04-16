package Go_LeetCode

import "math"

// 通过测试

func minCameraCover(root *TreeNode) int {
	data := process_968(root)
	return int(min_968(data.Uncovered+1, min_968(data.CoveredNoCamera, data.CoveredHasCamera)))
}

// 潜台词：x是头节点，x下方的点都被覆盖的情况下
type Info struct {
	Uncovered        int64 //x没有被覆盖，x为头的树至少需要几个相机
	CoveredNoCamera  int64 //x被相机覆盖，但是x没相机，x为头的树至少需要几个相机
	CoveredHasCamera int64 //x被相机覆盖了，并且x上放了相机，x为头的树至少需要几个相机
}

func process_968(X *TreeNode) Info {
	if X == nil {
		return Info{math.MaxInt32, 0, math.MaxInt32}
	}

	left := process_968(X.Left)
	right := process_968(X.Right)
	// x uncovered x自己不被覆盖，x下方所有节点，都被覆盖
	// 左孩子： 左孩子没被覆盖，左孩子以下的点都被覆盖
	// 左孩子被覆盖但没相机，左孩子以下的点都被覆盖
	// 左孩子被覆盖也有相机，左孩子以下的点都被覆盖
	uncovered := left.CoveredNoCamera + right.CoveredNoCamera
	// x下方的点都被covered，x也被cover，但x上没相机
	/*
		long coveredNoCamera = Math.min(
			// 1)
			left.coveredHasCamera + right.coveredHasCamera,

			Math.min(
					// 2)
					left.coveredHasCamera + right.coveredNoCamera,

					// 3)
					left.coveredNoCamera + right.coveredHasCamera));
	*/
	coveredNoCamera := min_968(left.CoveredHasCamera+right.CoveredHasCamera,
		min_968(left.CoveredHasCamera+right.CoveredNoCamera, left.CoveredNoCamera+right.CoveredHasCamera))
	// x下方的点都被covered，x也被cover，且x上有相机
	/*
		long coveredHasCamera =
			Math.min(left.uncovered, Math.min(left.coveredNoCamera, left.coveredHasCamera))

			+ Math.min(right.uncovered, Math.min(right.coveredNoCamera, right.coveredHasCamera))

			+ 1;
	*/
	coveredHasCamera := min_968(left.Uncovered, min_968(left.CoveredNoCamera, left.CoveredHasCamera)) +
		min_968(right.Uncovered, min_968(right.CoveredNoCamera, right.CoveredHasCamera)) +
		1
	return Info{uncovered, coveredNoCamera, coveredHasCamera}
}

func min_968(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
