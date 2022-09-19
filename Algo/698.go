package algo

func matchTarget(nums []int, target int) []int {
	for k := range nums {
		if nums[k] == target {
			return []int{nums[k]}
		} else if nums[k] < target {
			newNum := make([]int, len(nums))
			copy(newNum, nums)
			newNum = append(newNum[:k], newNum[k+1:]...)
			leftNums := matchTarget(newNum, target-nums[k])
			if leftNums != nil {
				return append(leftNums, nums[k])
			}
		}
	}
	return nil
}

func collection(nums []int, target, k int) bool {
	if k == 1 {
		for _, v := range nums {
			target -= v
		}
		return target == 0
	}
	for nums != nil && k != 0 {
		leftNums := matchTarget(nums, target)
		if leftNums == nil {
			return false
		}
		for _, v := range leftNums {
			for d := len(nums) - 1; d >= 0; d-- {
				if nums[d] == v {
					nums = append(nums[:d], nums[d+1:]...)
					break
				}
			}
		}
		k--
	}
	return nums == nil && k == 0
}

func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	return collection(nums, target, k)
}
