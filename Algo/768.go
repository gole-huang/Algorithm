package Algo

func addNum(array [][]int, num int) [][]int {
	for i := 0; i < len(array); i++ {
		if num < array[i][0] {
			array[len(array)-1][1] = 1
			return append(array[:i], array[len(array)-1])
		}
	}
	if num == array[len(array)-1][0] {
		array[len(array)-1][1]++
	} else {
		array = append(array, []int{num, 1})
	}
	return array
}

func MaxChunksToSorted(arr []int) int {
	res := [][]int{{arr[0], 1}}
	for i := 1; i < len(arr); i++ {
		res = addNum(res, arr[i])
	}
	sum := 0
	for _, v := range res {
		sum += v[1]
	}
	return sum
}
