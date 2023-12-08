package Go_LeetCode

//func Hanota(A []int, B []int, C []int) []int {
//	var N = len(A)
//	process(N, &A, &C, &B)
//	return C
//}
//
//func process(n int, from *[]int, to *[]int, help *[]int) {
//	if n < 0 {
//		return
//	}
//	if n == 1 {
//		*to = append(*to, (*from)[len(*from)-1])
//		*from = (*from)[:len(*from)-1]
//		return
//	}
//	process(n-1, from, help, to)
//	process(1, from, to, help)
//	process(n-1, help, to, from)
//}
