package algo

func BinSearch(nums []int, target int) int {
	head, tail := 0, len(nums)
	for head <= tail {
		mid := head + (tail-head)>>1
		if nums[mid] < target {
			tail = mid - 1
		} else if nums[mid] > target {
			head = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
