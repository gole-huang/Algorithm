package algo

func ConstructArray(n int, k int) []int {
	arr := make([]int, n)
	m := (k - 1) / 2
	i := 0
	for ; i < m*2; i += 2 {
		if i >= len(arr) {
			return nil
		}
		arr[i] = i/2 + 1
		arr[i+1] = n - i/2
	}
	for ; i < n; i++ {
		arr[i] = i + 1 - m
	}
	if k%2 == 0 {
		arr[n-1], arr[n-2] = arr[n-2], arr[n-1]
	}
	return arr
}
