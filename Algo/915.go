package algo

import "fmt"

func PartitionDisjoint(nums []int) int {
	left := [][2]int{{0, nums[0]}}
	for i := 1; i < len(nums); i++ {
		if nums[i] > left[len(left)-1][1] {
			left = append(left, [2]int{i, nums[i]})
		}
	}
	right := [][2]int{{len(nums) - 1, nums[len(nums)-1]}}
	for d := len(nums) - 2; d > -1; d-- {
		if nums[d] < right[len(right)-1][1] {
			right = append(right, [2]int{d, nums[d]})
		}
	}
	for lk := range left {
		for rk := range right {
			fmt.Println(left[lk], right[rk])
			if left[lk] == right[rk] {
				return right[rk][0] + 1
			} else if rk < len(right)-1 && left[lk][1] <= right[rk][1] && left[lk][1] > right[rk+1][1] && left[lk+1][0] > right[rk+1][0] {
				return right[rk+1][0] + 1
			}
		}
	}
	return len(nums) - 1
}
