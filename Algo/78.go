package algo

func Subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, nums}
	} else {
		arr := Subsets(nums[1:])
		res := make([][]int, 0)
		for _, v := range arr {
			res = append(res, v)
			k := make([]int, len(v)+1)
			copy(k, v)
			k[len(k)-1] = nums[0]
			res = append(res, k)
		}
		return res
	}
}
