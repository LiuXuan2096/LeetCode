package Go_LeetCode

import "sort"

type Envelope struct {
	w int
	h int
}

type Envelopes []*Envelope

func (e Envelopes) Len() int {
	return len(e)
}

func (e Envelopes) Less(i, j int) bool {
	if e[i].w != e[j].w {
		return e[i].w < e[j].w
	}
	return e[i].h > e[j].h
}

func (e Envelopes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	enlops := make(Envelopes, n)
	for i := 0; i < n; i++ {
		enlops[i] = &Envelope{
			envelopes[i][0],
			envelopes[i][1],
		}
	}

	sort.Sort(enlops)

	ends := make([]int, n)
	ends[0] = enlops[0].h
	right := 0

	for i := 1; i < n; i++ {
		l, r := 0, right
		for l <= r {
			m := (l + r) / 2
			if enlops[i].h > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		right = max_354(l, right)
		ends[l] = enlops[i].h
	}

	return right + 1
}

func max_354(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
