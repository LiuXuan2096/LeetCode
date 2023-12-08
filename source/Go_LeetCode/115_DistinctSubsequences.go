package Go_LeetCode

//func NumDistinct(s string, t string) int {
//	var source []byte = []byte(s)
//	// var template []byte = []byte(t)
//	count := 0
//	process(source, 0, "", &count, t)
//	return count % (10e9 + 7)
//}
//
//// str固定参数
//// 来到了str【index】字符 index是位置
//// str【0...index-1】已经走过了，之前的决定都在path上
////之前的决定已经不能改变了 就是path
//// str【index】还能决定
//func process(str []byte, index int, path string, count *int, pattern string) {
//	if index == len(str) {
//		if path == pattern {
//			*count++
//		}
//		return
//	}
//	// 没有要index位置的字符
//	process(str, index+1, path, count, pattern)
//	// 要了index位置的字符
//	path += string(str[index])
//	process(str, index+1, path, count, pattern)
//}
