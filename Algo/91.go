package algo

func fabonacci(n int) int {
	a, b, c := 1, 1, 2
	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return b
}
func NumDecodings(s string) int {
	if len(s) == 0 || len(s) == 1 && s != "0" {
		return 1
	} else if s[0] == '0' {
		return 0
	} else {
		res := 1
		for d := len(s) - 1; d > 0; d-- {
			switch {
			case s[d] == '0':
				if s[d-1] == '1' || s[d-1] == '2' {
					d--
					continue
				} else {
					return 0
				}
			case s[d-1] != '1' && s[d-1] != '2' || s[d] >= '7' && s[d-1] == '2':
				continue
			default:
				dd := 0
				for d >= dd+1 && (s[d-1-dd] == '1' || s[d-1-dd] == '2') {
					dd++
				}
				res *= fabonacci(dd)
				d -= dd
				if d >= 0 && s[d] <= '6' {
					d++
				}
			}
		}
		return res
	}
}
