package algo

func maxDiv(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
func NthMagicalNumber(n int, a int, b int) int {
	if a < b {
		a, b = b, a
	}
	div := maxDiv(a, b)
	divA, divB := a/div, b/div
	minMul := a * b / div
	mul := n / (divA + divB - 1)
	left := n % (divA + divB - 1)
	if left == 0 { //剩余步长为0，就是公倍数*公倍数步长
		return (minMul * mul) % 1000000007
	}
	ptrA, ptrB := 1, 1 //剩余步长不为0，使用双因子按1步长递进
	for i := 0; i < left-1; i++ {
		if a*ptrA < b*ptrB {
			ptrA++
		} else {
			ptrB++
		}
	}
	if a*ptrA < b*ptrB {
		return ((minMul*mul)%1000000007 + (a * ptrA)) % 1000000007
	} else {
		return ((minMul*mul)%1000000007 + (b * ptrB)) % 1000000007
	}
}
