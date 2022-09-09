package algo

func CanBeEqual(target []int, arr []int) bool {
	for len(target) > 0 {
		nomatch := true
		for j := 0; j < len(arr); j++ {
			if target[0] == arr[j] {
				target = target[1:]
				arr = append(arr[:j], arr[j+1:]...)
				nomatch = false
				break
			}
		}
		if nomatch {
			return false
		}
	}
	return true
}
