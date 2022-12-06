package algo

func minOperations(nums1 []int, nums2 []int) int {
	arr1, arr2 := [7]int{}, [7]int{}
	for _, v := range nums1 {
		arr1[0] += v
		arr1[v]++
	}
	for _, v := range nums2 {
		arr2[0] += v
		arr2[v]++
	}
	if arr1[0] == arr2[0] {
		return 0
	} else if arr1[0] > arr2[0] {
		arr1, arr2 = arr2, arr1
	}
	diff, cnt := arr2[0]-arr1[0], 0
	for k := 1; k < 6; k++ {
		minus := (6 - k) * (arr1[k] + arr2[7-k])
		if diff > minus {
			diff -= minus
			cnt += arr1[k] + arr2[7-k]
		} else {
			cnt += (diff + 5 - k) / (6 - k)
			return cnt
		}
	}
	return -1
}
