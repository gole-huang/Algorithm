package algo

func Search2(nums []int, target int) bool {
	head, tail, mid := 0, len(nums)-1, len(nums)>>1
	for nums[mid] != target && head < tail-1 {
		if nums[head] == target || nums[mid] == target || nums[tail] == target {
			return true
		} else if nums[head] == nums[mid] && nums[mid] == nums[tail] {
			if nums[(head+mid)>>1] != nums[mid] {
				tail = mid
			} else if nums[(mid+tail)>>1] != nums[mid] {
				head = mid
			} else {
				head++
				tail--
			}
		} else if nums[mid] < target { //target比mid大，位于队列右侧
			if target < nums[tail] || nums[tail] < nums[mid] {
				head = mid
			} else { //target比mid大，不位于队列右侧
				tail = mid
			}
		} else if target < nums[mid] { //target比mid小，位于队列左侧
			if nums[head] < target || nums[head] > nums[mid] {
				tail = mid
			} else { //target比mid小，不位于队列左侧
				head = mid
			}
		}
		mid = (head + tail) >> 1
	}
	return nums[head] == target || nums[mid] == target || nums[tail] == target
}
