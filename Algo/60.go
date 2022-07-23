package Algo

func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func GetPermutation(n int, k int) string {
	k--
	srl := make([]byte, n)
	for i := 0; i < n; i++ {
		srl[i] = byte(49 + i)
	}
	res := make([]byte, n)
	for i := 0; i < len(res)-1; i++ {
		nStep := factorial(n - 1)
		tmp := k / nStep
		res[i] = srl[tmp]
		srl = append(srl[:tmp], srl[tmp+1:]...)
		k %= nStep
		n--
	}
	res[len(res)-1] = srl[0]
	return string(res)
}
