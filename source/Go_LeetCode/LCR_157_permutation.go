package Go_LeetCode

//func GoodsOrder(goods string) []string {
//	var goodsArray []byte = []byte(goods)
//	var ans []string
//	process(&goodsArray, 0, &ans)
//	return ans
//}
//
//func process(str *[]byte, index int, ans *[]string) {
//	if index == len(*str) {
//		*ans = append(*ans, string(*str))
//	} else {
//		var visited [256]bool
//		for i := index; i < len(*str); i++ {
//			if !visited[(*str)[i]] {
//				visited[(*str)[i]] = true
//				swap(str, index, i)
//				process(str, index+1, ans)
//				swap(str, i, index)
//			}
//		}
//	}
//}
//
//func swap(chs *[]byte, i, j int) {
//	var tmp byte = (*chs)[i]
//	(*chs)[i] = (*chs)[j]
//	(*chs)[j] = tmp
//}
