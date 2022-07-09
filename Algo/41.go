package Algo

/*
func FirstMissingPositive(nums []int) int {
	head, tail := 0, len(nums)-1
	for head < tail { //让所有正整数在前，非正整数在后
		for nums[head] < len(nums) && nums[head] > 0 {
			head++
		}
		for nums[tail] >= len(nums) || nums[tail] <= 0 {
			tail--
		}
		if (nums[head] >= len(nums) || nums[head] <= 0) && (nums[tail] < len(nums) && nums[tail] > 0) {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		}
	}
	cnt := head //有cnt个
	newNum := nums[:tail]
}
*/
