package algo

func FirstMissingPositive(nums []int) int {
	//nums = append([]int{0}, nums...)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 1 && nums[i] <= len(nums) {
			if nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
				i--
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
