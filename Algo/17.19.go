package algo

func MissingTwo(nums []int) []int {
	n_1, n := 0, 0
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != 0 {
			if nums[i] == len(nums)+2 {
				nums[i], n = n, nums[i]
			} else if nums[i] == len(nums)+1 {
				nums[i], n_1 = n_1, nums[i]
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}
	res := make([]int, 0)
	for k := range nums {
		if nums[k] == 0 {
			res = append(res, k+1)
		}
	}
	if len(res) < 2 {
		if n_1 == 0 {
			res = append(res, len(nums)+1)
		}
		if n == 0 {
			res = append(res, len(nums)+2)
		}
	}
	return res
}
