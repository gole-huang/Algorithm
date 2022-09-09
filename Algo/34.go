package algo

import "fmt"

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	left, right := 0, len(nums)-1
	var head, tail int
	for left < right {
		mid := (left + right) >> 1
		if mid == left {
			break
		}
		fmt.Printf("Target: nums[%d]:%d\n", mid, nums[mid])
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else { //找到其中一个与目标相等的数字
			head, tail = mid, mid
			for left < head { //找左边界
				mid = (left + head) >> 1
				if mid == left {
					break
				}
				fmt.Printf("LeftBoundary: nums[%d]:%d\n", mid, nums[mid])
				if nums[mid] == target { //只要中值等于目标，则左边界在中值左侧
					head = mid
				} else { //中值小于目标，则左边界在右侧
					left = mid
				}
			}
			for tail < right { //找右边界
				mid = (tail + right) >> 1
				if mid == tail {
					break
				}
				fmt.Printf("RightBoundary: [%d]:%d\n", mid, nums[mid])
				if nums[mid] == target { //只要中值等于目标，则右边界在中值右侧
					tail = mid
				} else { //中值小于目标，则左边界在右侧
					right = mid
				}
			}
			break
		}
	}
	if nums[left] == target && nums[right] == target {
		return []int{left, right}
	} else if nums[left] == target {
		if nums[tail] == target {
			return []int{left, tail}
		} else {
			return []int{left, left}
		}
	} else if nums[right] == target {
		if nums[head] == target {
			return []int{head, right}
		} else {
			return []int{right, right}
		}
	} else if left == right-1 {
		return []int{-1, -1}
	} else {
		return []int{head, tail}
	}
}
