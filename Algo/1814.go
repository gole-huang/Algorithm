package algo

func countNicePairs(nums []int) (res int) {
	cnt := make(map[int]int)
	for k, v := range nums {
		key := func() int {
			tmp := 0
			for v != 0 {
				tmp = tmp*10 + v%10
				v /= 10
			}
			return nums[k] - tmp
		}()
		cnt[key]++
	}
	for _, v := range cnt {
		res += v * (v - 1) >> 1
		res %= 1000000007
	}
	return
}
