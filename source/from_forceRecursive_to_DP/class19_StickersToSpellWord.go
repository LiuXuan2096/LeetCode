package from_forceRecursive_to_DP

import "math"

/*
题目测试连接：https://leetcode.com/problems/stickers-to-spell-word
给定一个字符串str，给定一个字符串类型的数组arr，出现的字符都是小写英文
arr每一个字符串，代表一张贴纸，你可以把单个字符剪开使用，目的是拼出str来
返回需要至少多少张贴纸可以完成这个任务
例子：str= "babac"，arr = {"ba","c","abcd"}
ba + ba + c  3  abcd + abcd 2  abcd+ba 2
所以返回2
*/

// MinStacker 暴力递归版本的解法
func MinStacker(stickers []string, target string) int {
	N := len(stickers)
	// 用词频代替贴纸数组
	counts := make([][]int, N)
	for i := 0; i < N; i++ {
		counts[i] = make([]int, 26)
		str := []rune(stickers[i])
		for _, ch := range str {
			counts[i][ch-'a']++
		}
	}
	ans := process_C19_3(counts, target)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// stickers[i] 数组，当初i号贴纸的字符统计 [][]int stickers -> 所有的贴纸
// 每一种贴纸都有无穷张
// 返回搞定target的最少张数
// 最少张数
func process_C19_3(stickers [][]int, t string) int {
	if len(t) == 0 {
		return 0
	}
	target := []rune(t)
	tcounts := make([]int, 26)
	for _, ch := range target {
		tcounts[ch-'a']++
	}
	N := len(stickers)
	res := math.MaxInt32
	for i := 0; i < N; i++ {
		sticker := stickers[i]          // 尝试第一张贴纸
		if sticker[target[0]-'a'] > 0 { // 剪枝
			var builder []rune
			for j := 0; j < 26; j++ {
				if tcounts[j] > 0 {
					// nums:用了当前这张贴纸后，还剩下几个当前字符
					nums := tcounts[j] - sticker[j]
					for k := 0; k < nums; k++ {
						builder = append(builder, rune(j+'a'))
					}
				}
			}
			rest := string(builder)
			res = min(res, process_C19_3(stickers, rest))
		}
	}
	if res == math.MaxInt32 {
		return res
	} else {
		return res + 1
	}
}

// MinSticker3 傻缓存版本
func MinSticker3(stickers []string, target string) int {
	N := len(stickers)
	counts := make([][]int, N)
	for i := 0; i < N; i++ {
		counts[i] = make([]int, 26)
		str := []rune(stickers[i])
		for _, ch := range str {
			counts[i][ch-'a']++
		}
	}

	dp := make(map[string]int)
	dp[""] = 0

	ans := process_C19_3_2(counts, target, dp)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func process_C19_3_2(stickers [][]int, t string, dp map[string]int) int {
	if val, ok := dp[t]; ok {
		return val
	}

	target := []rune(t)
	tcounts := make([]int, 26)
	for _, ch := range target {
		tcounts[ch-'a']++
	}

	N := len(stickers)
	res := math.MaxInt32
	for i := 0; i < N; i++ {
		sticker := stickers[i]          // 取出一张贴纸
		if sticker[target[0]-'a'] > 0 { // 剪枝
			var builder []rune
			for j := 0; j < 26; j++ {
				if tcounts[j] > 0 {
					nums := tcounts[j] - sticker[j]
					for k := 0; k < nums; k++ {
						builder = append(builder, rune(j+'a'))
					}
				}
			}
			rest := string(builder)
			res = min(res, process_C19_3_2(stickers, rest, dp))
		}
	}

	var ans int
	if res == math.MaxInt32 {
		ans = res
	} else {
		ans = res + 1
	}
	dp[t] = ans
	return ans
}
