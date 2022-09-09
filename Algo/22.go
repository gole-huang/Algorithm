package algo

func addParentheses(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	var ss []string
	for _, s := range addParentheses(n - 1) {
		ss = append(ss, "()"+s)
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				ss = append(ss, "("+s[:i+1]+")"+s[i+1:])
			}
		}
	}
	return ss
}

func GenerateParenthesis(n int) []string {
	str := addParentheses(n)
	res := make([]string, 0)
	strMap := make(map[string]bool)
	for _, s := range str {
		_, ok := strMap[s]
		if !ok {
			res = append(res, s)
			strMap[s] = true
		}
	}
	return res
}
