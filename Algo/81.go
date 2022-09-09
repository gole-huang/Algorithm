package algo

func normalSearch(nums []int, target, left, right int) bool {
	if nums[left] > target || nums[right] < target {
		return false
	}
	for left < right-1 {
		if nums[left] == nums[right] {
			return nums[left] == target
		}
		mid := (left + right) >> 1
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[left] == target || nums[right] == target
}

func Newsearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	if target == nums[left] || target == nums[right] {
		return true
	}
	for left < right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			return true
		} else if target < nums[mid] {
			if target == nums[left] {
				return true
			} else if target > nums[left] {
				return normalSearch(nums, target, left, mid)
			} else { //if target < nums[left]
				if nums[left] == nums[mid] {
					for nums[left] == nums[left+1] && left < mid {
						left++
					}
					for nums[mid] == nums[mid-1] && left < mid {
						mid--
					}
					if left == mid {
						continue
					}
				}
				if nums[left] > nums[mid] {
					right = mid
					continue
				} else { //nums[left] < nums[mid]
					left = mid
					continue
				}
			}
		} else { //target > nums[mid]
			if target == nums[right] {
				return true
			} else if target < nums[right] {
				return normalSearch(nums, target, mid, right)
			} else { //if target > nums[right]
				if nums[mid] == nums[right] {
					for nums[mid] == nums[mid+1] && mid < right {
						mid++
					}
					for nums[right] == nums[right-1] && mid < right {
						right--
					}
					if right == mid {
						continue
					}
				}
				if nums[mid] > nums[right] {
					right = mid
					continue
				} else { //nums[mid] < nums[right]
					left = mid
					continue
				}
			}
		}
	}
	return false
}
