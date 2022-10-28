package algo

// LetterCasePermutation '0'-'9':48-57	'A'-'Z':65-90	'a'-'z':97-122
func LetterCasePermutation(s string) []string {
	if len(s) == 1 {
		switch {
		case s[0] < 61:
			return []string{s}
		case s[0] > 94:
			return []string{string(s[0] - 32), s}
		default:
			return []string{s, string(s[0] + 32)}
		}
	} else {
		lgh := len(s)
		str := LetterCasePermutation(s[:lgh-1])
		switch {
		case s[lgh-1] < 61:
			for k := range str {
				str[k] += string(s[lgh-1])
			}
		case s[lgh-1] > 94:
			strLen := len(str)
			str = append(str, str...)
			for k := 0; k < strLen; k++ {
				str[k] += string(s[lgh-1] - 32)
				str[k+strLen] += string(s[lgh-1])
			}
		default:
			strLen := len(str)
			str = append(str, str...)
			for k := 0; k < strLen; k++ {
				str[k] += string(s[lgh-1])
				str[k+strLen] += string(s[lgh-1] + 32)
			}
		}
		return str
	}
}
