package algo

import "sort"

func NextPermutation(nums []int) {
	lgh := len(nums) - 1
	if lgh == 0 {
		return
	}
	i := lgh
	for ; i > 0; i-- { //找出正序位置
		if nums[i] > nums[i-1] {
			if i == lgh {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				return
			}
			j := i + 1 //找出比nums[i-1]大的最小值
			for ; j <= lgh; j++ {
				if nums[j] <= nums[i-1] {
					break
				}
			}
			if nums[j-1] > nums[i-1] {
				nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
			} else {
				nums[lgh], nums[i-1] = nums[i-1], nums[lgh]
			}
			break
		}
	}
	if i == 0 { //nums为倒序数组，直接返回正序排序
		sort.Ints(nums)
	} else { //对从第i个往后的数进行排序
		sort.Ints(nums[i:])
	}
	return
}
