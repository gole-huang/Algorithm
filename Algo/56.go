package Algo

import (
	"sort"
)

type Arr2 [][]int

func (a Arr2) Len() int {
	return len(a)
}
func (a Arr2) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a Arr2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Sort(Arr2(intervals))
	tmp := []int{intervals[0][0], intervals[0][1]}
	intervals = intervals[1:]
	for len(intervals) != 0 {
		if tmp[1] < intervals[0][0] {
			res = append(res, tmp)
			tmp = []int{intervals[0][0], intervals[0][1]}
		} else if tmp[1] < intervals[0][1] {
			tmp[1] = intervals[0][1]
		}
		intervals = intervals[1:]
	}
	res = append(res, tmp)
	return res
}
