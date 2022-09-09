package algo

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	val := 1
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-1-i; j++ {
			res[i][j] = val + j - i
			res[j][n-1-i] = val + n - 1 - 3*i + j
			res[n-1-i][n-1-j] = val + 2*n - 2 - 5*i + j
			res[n-1-j][i] = val + 3*n - 3 - 7*i + j
		}
		val += 4 * (n - 2*i - 1)
	}
	if n%2 != 0 {
		res[n/2][n/2] = n * n
	}
	return res
}
