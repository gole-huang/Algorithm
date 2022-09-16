package algo

import "sort"

func addSegment(lines [][2]int, head, tail int) [][2]int {
	switch {
	case len(lines) == 0:
		return [][2]int{{head, tail}}
	case tail < lines[0][0]:
		return append([][2]int{{head, tail}}, lines...)
	case head > lines[len(lines)-1][1]:
		return append(lines, [2]int{head, tail})
	default:
		tailEnd := true
		headSegment, tailSegment := 0, 0
		for k := range lines {
			if head < lines[k][1] {
				headSegment = k
				if head > lines[headSegment][0] {
					head = lines[headSegment][0]
				}
				break
			}
		}
		for k := headSegment; k < len(lines); k++ {
			if tail < lines[k][0] {
				tailEnd = false
				tailSegment = k - 1
				if tail < lines[tailSegment][1] {
					tail = lines[tailSegment][1]
				}
				break
			}
		}
		if tailEnd {
			if tail < lines[len(lines)-1][1] {
				tail = lines[len(lines)-1][1]
			}
			return append(lines[:headSegment], [2]int{head, tail})
		}
		tmpLines := make([][2]int, len(lines[tailSegment+1:]))
		copy(tmpLines, lines[tailSegment+1:])
		lines = append(lines[:headSegment], [2]int{head, tail})
		lines = append(lines, tmpLines...)
		return lines
	}
}

func RectangleArea(rectangles [][]int) int {
	if len(rectangles) == 1 {
		return int(int64((rectangles[0][2]-rectangles[0][0])*(rectangles[0][3]-rectangles[0][1])) % 1000000007)
	}
	sort.Slice(rectangles, func(i, j int) bool {
		switch {
		case rectangles[i][0] != rectangles[j][0]:
			return rectangles[i][0] < rectangles[j][0]
		case rectangles[i][2] != rectangles[j][2]:
			return rectangles[i][2] < rectangles[j][2]
		case rectangles[i][1] != rectangles[j][1]:
			return rectangles[i][1] < rectangles[j][1]
		default:
			return rectangles[i][3] < rectangles[j][3]
		}
	})
	res := int64(0)
	xMap := make(map[int]bool, 0)
	for rKey, rVal := range rectangles {
		for x := rVal[0]; x <= rVal[2]; x++ {
			if _, ok := xMap[x]; !ok {
				xMap[x] = true
				lines := make([][2]int, 0)
				for _, rrVal := range rectangles[rKey:] {
					if x < rrVal[0] {
						break
					} else if x < rrVal[2] {
						lines = addSegment(lines, rrVal[1], rrVal[3])
					}
				}
				for _, v := range lines {
					res += int64(v[1] - v[0])
				}
			}
		}
	}
	return int(res % 1000000007)
}
