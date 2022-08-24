package Algo

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func FindClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	left, right := 0, len(arr)-1
	var mid int
	for left < right-1 {
		mid = (left + right) >> 1
		if x < arr[mid] {
			right = mid
		} else if x > arr[mid] {
			left = mid
		} else {
			break
		}
	}
	res, t := 0, k-1
	switch {
	case x == arr[mid]:
		res = mid
	case abs(x, arr[left]) <= abs(x, arr[right]):
		res = left
	default:
		res = right
	}
	left, right = res-1, res+1
	for ; t > 0; t-- {
		if left == -1 {
			res = 0
			break
		}
		if right == len(arr) || abs(x, arr[left]) <= abs(x, arr[right]) {
			res = left
			left--
		} else {
			res = right
			right++
		}
	}
	return arr[left+1 : left+1+k]
}
