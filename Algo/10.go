package Algo

import "fmt"

func anyChar(s, p string) bool {
	//p当前匹配字符为.*时调用
	fmt.Printf("Match .*\n")
	if p == "" {
		fmt.Printf("P empty\n")
		return true
	} else {
		for i := 0; i < len(s); i++ {
			fmt.Printf("Match start from : %q\n", s[i])
			if IsMatch(s[i:], p) {
				return true
			}
		}
		return IsMatch("", p)
	}
}

func sameChar(s, p string, c byte) bool {
	fmt.Printf("Matching %c*\n", c)
	for j := 0; j < len(s) && s[j] == c; j++ {
		fmt.Printf("Match %s with %s\n", s[j:], p)
		if IsMatch(s[j+1:], p) {
			return true
		}
	}
	return IsMatch(s, p)
}

func IsMatch1(s string, p string) bool {
	fmt.Printf("s:%s , p:%s\n", s, p)
	i, j := 0, 0
	for i < len(p) {
		if i == len(p)-1 {
			fmt.Printf("P Last character : %q\n", p[i])
			return j == len(s)-1 && (p[i] == s[j] || p[i] == '.')
		} else if j == len(s) {
			fmt.Printf("P left %s\n", p[i:])
			for i < len(p)-1 {
				if p[i+1] == '*' {
					fmt.Printf("p[%d] is '*', pass 2 characters\n", i+1)
					i += 2
				} else {
					return false
				}
			}
		} else if p[i] == '.' {
			if p[i+1] != '*' {
				fmt.Printf("p[%d] = '.' , s[%d] = %q\n", i, j, s[j])
				i++
				j++
			} else {
				return anyChar(s[j:], p[i+2:])
			}
		} else if p[i+1] == '*' {
			fmt.Printf("p[%d:%d] = '%c*', s[%d] = %q\n", i, i+1, p[i], j, s[j])
			return sameChar(s[j:], p[i+2:], p[i])
		} else if p[i] == s[j] {
			fmt.Printf("Match, s[%d] = p[%d] = %q\n", j, i, p[i])
			i++
			j++
		} else {
			return false
		}
	}
	return j == len(s)
}
