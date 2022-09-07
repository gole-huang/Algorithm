package Algo

func ValidateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	} else if len(pushed) == 0 {
		return true
	} else if len(pushed) == 1 {
		return pushed[0] == popped[0]
	}
	pos := 0
	for ; pos < len(popped); pos++ {
		if popped[pos] == pushed[0] {
			break
		}
	}
	if pos == len(popped) {
		return false
	}
	return ValidateStackSequences(pushed[1:pos+1], popped[:pos]) && ValidateStackSequences(pushed[pos+1:], popped[pos+1:])
}
