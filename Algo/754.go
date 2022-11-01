package algo

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	n, i := 0, 0
	for n < target {
		i++
		n += i
	}
	if (n-target)%2 == 0 {
		return i
	} else if i%2 == 0 {
		return i + 1
	} else {
		return i + 2
	}
}
