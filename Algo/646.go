package algo

import "sort"

type arrWarpper [][]int

func (arr arrWarpper) Len() int {
	return len(arr)
}

func (arr arrWarpper) Less(a, b int) bool {
	return arr[a][1] < arr[b][1]
}

func (arr arrWarpper) Swap(a, b int) {
	arr[a][0], arr[b][0] = arr[b][0], arr[a][0]
	arr[a][1], arr[b][1] = arr[b][1], arr[a][1]
}

func FindLongestChain(pairs [][]int) int {
	sort.Sort(arrWarpper(pairs))
	cnt := 1
	tail := pairs[0][1]
	for _, v := range pairs {
		if v[0] > tail {
			cnt++
			tail = v[1]
		}
	}
	return cnt
}
