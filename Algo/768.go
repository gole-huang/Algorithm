package Algo

func addNum(array []int, num int) []int {
	for i := 0; i < len(array); i++ {
		if num < array[i] {
			return append(array[:i], array[len(array)-1])
		}
	}
	return append(array, num)
}

func maxChunksToSorted(arr []int) int {
	res := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		res = addNum(res, arr[i])
	}
	return len(res)
}
