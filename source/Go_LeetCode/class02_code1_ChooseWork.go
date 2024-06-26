package Go_LeetCode

/*
给定数组hard和money，长度都为N，hard[i]表示i号工作的难度，
money[i]表示i号工作的收入给定数组ability，长度都为M，
ability[j]表示j号人的能力，每一号工作，都可以提供无数的岗位，
难度和收入都一样
但是人的能力必须>=这份工作的难度，才能上班。
返回一个长度为M的数组ans，ans[j]表示j号人能获得的最好收入
*/

//type Job struct {
//	money int
//	hard  int
//}
//
//type JobComparator []Job
//
//func (c JobComparator) Len() int {
//	return len(c)
//}
//
//func (c JobComparator) Less(i, j int) bool {
//	if c[i].hard != c[j].hard {
//		return c[i].hard < c[j].hard
//	} else {
//		return c[j].money < c[i].money
//	}
//}
//
//func (c JobComparator) Swap(i, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//func GetMoneys(jobs []Job, ability []int) int {
//	sort.Sort(JobComparator(jobs))
//
//	// key :难度 value: 报酬
//	mapJob := make(map[int]int)
//	mapJob[jobs[0].hard] = jobs[0].money
//}
