package algo

func LargestRectangleArea(heights []int) int {
	maxSqr := 0
	curSqrs := make([][2]int, 0)
	for _, v := range heights {
		if len(curSqrs) == 0 {
			curSqrs = [][2]int{{v, 1}}
			continue
		}
		maxLength, noExist := 0, true
		for i := 0; i < len(curSqrs); i++ {
			if curSqrs[i][0] > v {
				if curSqrs[i][0]*curSqrs[i][1] > maxSqr {
					maxSqr = curSqrs[i][0] * curSqrs[i][1]
				}
				if curSqrs[i][1] > maxLength {
					maxLength = curSqrs[i][1]
				}
				curSqrs = append(curSqrs[:i], curSqrs[i+1:]...)
				i--
			} else {
				if curSqrs[i][0] == v {
					noExist = false
				}
				curSqrs[i][1]++
			}
		}
		if noExist && v != 0 {
			curSqrs = append(curSqrs, [2]int{v, maxLength + 1})
		}
	}
	for _, v := range curSqrs {
		if v[0]*v[1] > maxSqr {
			maxSqr = v[0] * v[1]
		}
	}
	return maxSqr
}
