package algo

func NumSubarrayBoundedMax(nums []int, left int, right int) int {
	res := 0
	edge, peak := -1, -1
	for k := range nums {
		if nums[k] > right {
			if edge != k-1 && peak != -1 {
				res += (k - 1 - peak) * (peak - edge)
			}
			edge, peak = k, -1
		} else {
			if nums[k] >= left && nums[k] <= right {
				if peak != -1 {
					res += (k - peak) * (peak - edge + 1)
				} else {
					res += k - edge
				}
				peak = k
			}
		}
	}
	if peak != -1 {
		res += (len(nums) - 1 - peak) * (peak - edge)
	}
	return res
}
