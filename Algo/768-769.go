package algo

func MaxChunksToSorted(arr []int) int {
	res := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		j := 0
		for ; j < len(res); j++ {
			if res[j] > arr[i] {
				res = append(res[:j], res[len(res)-1])
				break
			}
		}
		if j == len(res) {
			res = append(res, arr[i])
		}
	}
	return len(res)
}
