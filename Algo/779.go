package algo

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	} else {
		if k <= 1<<(n-2) {
			return kthGrammar(n-1, k)
		}
		return 1 - kthGrammar(n-1, k-1<<(n-2))
	}
}
