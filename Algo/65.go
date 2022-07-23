package Algo

func IsNumber(s string) bool {
	hasDot, hasE, hasNum := false, false, false
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if i == 0 || s[i-1] == 'e' || s[i-1] == 'E' {
				continue
			} else {
				return false
			}
		} else if s[i] == '.' {
			if !hasDot && !hasE {
				hasDot = true
			} else {
				return false
			}
		} else if s[i] == 'e' || s[i] == 'E' {	//E前面必须是一个非零数，后面必须是数
			if !hasE && hasNum {
				hasE = true
				hasNum = false
			} else {
				return false
			}
		} else if s[i] == 48 {
			hasNum = true
			continue
		} else if s[i] > 48 && s[i] < 58 {
			hasNum = true
			continue
		} else {
			return false
		}
	}
	if hasDot && !hasNum {
		return false
	} else {
		return hasNum
	}
}
