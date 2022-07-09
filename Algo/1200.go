package Algo

import (
	"sort"
)

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min := arr[1] - arr[0]
	res := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = make([]int, 2)
	}
	res[0][0], res[0][1] = arr[0], arr[1]
	j := 1
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < min {
			min, j = arr[i+1]-arr[i], 1
			res[0][0], res[0][1] = arr[i], arr[i+1]
		} else if arr[i+1]-arr[i] == min {
			res[j][0], res[j][1] = arr[i], arr[i+1]
			j++
		}

	}
	return res[:j]
}
