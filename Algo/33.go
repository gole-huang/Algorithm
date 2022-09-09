package algo

func divSearch(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head < tail-1 {
		mid := (head + tail) >> 1
		if nums[mid] < target {
			head = mid
		} else if nums[mid] > target {
			tail = mid
		} else {
			return mid
		}
	}
	if nums[head] == target {
		return head
	} else if nums[tail] == target {
		return tail
	} else {
		return -1
	}
}

func Search(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head < tail-1 {
		mid := (head + tail) >> 1
		if nums[mid] < target {
			if nums[tail] == target {
				return tail
			} else if nums[tail] > target { //在正常的序列中
				res := BinSearch(nums[mid:], target)
				if res >= 0 {
					return res + mid
				} else {
					return res
				}
			} else { //在包含k的序列中
				if nums[mid] > nums[tail] {
					head = mid
				} else {
					tail = mid
				}
			}
		} else if nums[mid] > target {
			if nums[head] == target {
				return head
			} else if nums[head] < target { //在正常的序列中
				return BinSearch(nums[:mid], target)
			} else { //在包含k的序列中head = mid
				if nums[mid] > nums[head] {
					head = mid
				} else {
					tail = mid
				}
			}
		} else {
			return mid
		}
	}
	if target == nums[head] {
		return head
	} else if target == nums[tail] {
		return tail
	} else {
		return -1
	}
}
