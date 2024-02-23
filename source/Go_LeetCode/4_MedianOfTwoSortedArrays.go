package Go_LeetCode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	even := size%2 == 0

	if len(nums1) != 0 && len(nums2) != 0 {
		if even {
			// 两数组长度相加是偶数，
			return float64(findKthNum(nums1, nums2, size/2)+findKthNum(nums1, nums2, size/2+1)) / 2.0
		} else {
			// 两数组长度相加是奇数
			return float64(findKthNum(nums1, nums2, size/2+1))
		}
	} else if len(nums1) != 0 {
		if even {
			return float64(nums1[size/2]+nums1[size/2-1]) / 2.0
		} else {
			return float64(nums1[size/2])
		}
	} else if len(nums2) != 0 {
		if even {
			return float64(nums2[size/2]+nums2[size/2-1]) / 2.0
		} else {
			return float64(nums2[size/2])
		}
	} else {
		return 0
	}
}

// findKthNum 找到两个正序数组中 第k小的数
func findKthNum(arr1, arr2 []int, kth int) int {
	// 先确定较长的数组和较短的数组
	longs := arr1
	shorts := arr2
	if len(longs) < len(shorts) {
		longs = arr2
		shorts = arr1
	}
	l := len(longs)
	s := len(shorts)

	// 如果 kth 小于等于 s 则第k小的数出现在 shorts 和 longs中从0到kth-1这s个数组成的新数组中
	// 并且这个第k小的数 就是这两个数组的上中位数
	if kth <= s {
		return getUpMedian(shorts, 0, kth-1, longs, 0, kth-1)
	}

	if kth > l {
		if shorts[kth-l-1] >= longs[l-1] {
			// 如果s数组的 第kth-l个数 比l数组的最后一个数还大
			// 说明s数组的 第kth-l个数 就是两个数组中第k小的数
			return shorts[kth-l-1]
		}
		if longs[kth-s-1] >= shorts[s-1] {
			// 这个if语句的道理和上一个一样
			return longs[kth-s-1]
		}
		// shorts数组中的前 kth-l 个数不参与计算上中位数
		// longs数组中的前 kth-s 个数不参与计算上中位数
		// 参与计算上中位数的区间长度是 s+l-kth
		// 假设所得的上中位数为 median 则 在参与计算上中位数的两个数组中
		// 有s+l-kth-1个数比median小
		// 所以两个数组总计有 kth-1个数比median小
		// 所以median就是两个数组中第k小的数
		return getUpMedian(shorts, kth-l, s-1, longs, kth-s, l-1)
	}

	// 执行到这里 说明 s<kth<=l
	if longs[kth-s-1] >= shorts[s-1] {
		// 小标为kth-s-1的数前面有kth-s-1个数比它小
		// 如果这个数再比shorts的最后一个数大 shorts中有s个数
		// 那么两个数组中共计有 kth-1 个数比它小
		// 这样 它就是 第k小的数
		return longs[kth-s-1]
	}
	// 在s<kth<=l情况下：
	// 长数组的前kth-s个数不参与求上中位数，短数组全部参与
	// 这样参与求上中位数的区间长度是s，假设求得的上中位数是median
	// 则在参与求上中位数的区间中 有s-1个数比它小，
	// s-1 + kth -s = kth-1 这样在两个数组中共计有kth-1个数比它小
	// 这样median就是第k小的数
	return getUpMedian(shorts, 0, s-1, longs, kth-s, kth-1)
}

// getUpMedian 返回两个长度相同的正序数组的上中位数
func getUpMedian(A []int, s1, e1 int, B []int, s2, e2 int) int {
	var mid1, mid2 int

	for s1 < e1 {
		mid1 = (s1 + e1) / 2
		mid2 = (s2 + e2) / 2 // 先求出两个数组的中间位置
		if A[mid1] == B[mid2] {
			return A[mid1] // 两个数组的中位数相等，则这个数就是两个数组共同的上中位数
		}

		if (e1-s1+1)%2 == 1 {
			// 说明数组长度是奇数
			// 假设数组A为[1,2,3,4,5] 数组B为[1,2,3,4,5]
			// 数组中的值表示位置，不表示值，即表示第i个数的意思
			// 如果A[3] > B[3] 则两个数组中可能出现答案的位置是
			// A数组中的[1,2] 和 B数组中的[3,4,5]并且这两个区间的
			// 上中位数就是这两个数组的上中位数，这是个结论，我不会证明
			if A[mid1] > B[mid2] {
				if B[mid2] >= A[mid1-1] {
					// 这里手动判断 B[3]是否是这个区间的上中位数
					// 能进入这个if语句说明B[3]是这个区间的上中位数 返回B[3]
					// 手动判断的原因是保证 求上中位数的区间长度一样
					// 如果B[3]不是答案，扔掉B[3]
					// 继续求A[1,2] 和 B[4,5]区间上的上中位数
					return B[mid2]
				}
				e1 = mid1 - 1
				s2 = mid2 + 1
			} else {
				// 和上一种情况道理一样
				if A[mid1] >= B[mid2-1] {
					return A[mid1]
				}
				e2 = mid2 - 1
				s1 = mid1 + 1
			}
		} else {
			// 执行到这里说明数组长度是偶数
			// 假设A[1,2,3,4,5,6] B[1,2,3,4,5,6]
			if A[mid1] > B[mid2] {
				// 这种情况下 上中位数出现的区间是A[1,2,3]和B[1,2,3]
				// 而且答案就是这两个区间的上中位数，这是个结论，不会证明
				e1 = mid1
				s2 = mid2 + 1
			} else {
				s1 = mid1 + 1
				e2 = mid2
			}
		}
	}
	return min(A[s1], B[s2])
}

//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
