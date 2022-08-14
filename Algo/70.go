package Algo

func ClimbStairs(n int) int {
	pre, cur := 1, 1
	for i := 0; i < n-1; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
