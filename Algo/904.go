package algo

func TotalFruit(fruits []int) int {
	num1, num2 := fruits[0], -1
	cnt1, cnt2, maxFruit := 1, 0, 0
	for i := 1; i < len(fruits); i++ {
		switch {
		case num1 == fruits[i]:
			cnt1++
		case num2 == fruits[i]:
			cnt2++
		case num2 == -1:
			num2 = fruits[i]
			cnt2 = 1
		default:
			if maxFruit < cnt1+cnt2 {
				maxFruit = cnt1 + cnt2
			}
			last := 1
			for d := i - 1; d > 0; d-- {
				if fruits[d-1] != fruits[d] {
					break
				} else {
					last++
				}
			}
			if fruits[i-1] == num2 {
				num1 = fruits[i]
				cnt1 = 1
				cnt2 = last
			} else {
				cnt1 = last
				num2 = fruits[i]
				cnt2 = 1
			}
		}
	}
	if maxFruit < cnt1+cnt2 {
		return cnt1 + cnt2
	}
	return maxFruit
}
