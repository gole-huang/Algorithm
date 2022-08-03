package Algo

import "sort"

type byteSort []byte

func (b byteSort) Len() int {
	return len(b)
}
func (b byteSort) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byteSort) Less(i, j int) bool {
	return b[i] < b[j]
}
func compareSubString(c1, c2 int, c []byte) int {
	for i := 1; i < len(c); i++ {
		if c[(c1+i)%len(c)] == c[(c2+i)%len(c)] {
			continue
		} else if c[(c1+i)%len(c)] < c[(c2+i)%len(c)] {
			return c1
		} else {
			return c2
		}
	}
	return c1
}
func OrderlyQueue(s string, k int) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	c := []byte(s)
	if k == 1 {
		min := c[0]
		pos := make([]int, 0)
		for i := 1; i < len(c); i++ {
			if c[i] < min {
				min = c[i]
			}
		}
		for i := 0; i < len(c); i++ {
			if c[i] == min {
				pos = append(pos, i)
			}
		}
		head := pos[0]
		for i := 1; i < len(pos); i++ {
			head = compareSubString(head, pos[i], c)
		}
		return string(append(c[head:], c[:head]...))
	} else {
		sort.Sort(byteSort(c))
		return string(c)
	}
}
