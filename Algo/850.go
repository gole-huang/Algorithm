package algo

import "sort"

type lineSegment struct {
	no      int
	segment [][2]int
	length  int
}

func addLine(lines [][2]int, head, tail int) [][2]int {
	if tail < lines[0][0] {
		return append([][2]int{{head, tail}}, lines...)
	} else if head > lines[len(lines)-1][1] {
		return append(lines, [2]int{head, tail})
	} else {
		headSegment, tailSegment := 0, 0
		for k := range
	}
}

func rectangleArea(rectangles [][]int) int {
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
	xMap := make(map[int]int, 0)
	for rKey, rVal := range rectangles {
		for x := rVal[0]; x <= rVal[2]; x++ {
			if _, ok := xMap[x]; !ok {
				lines := []lineSegment{{x, [][2]int{{rVal[1], rVal[3]}}, 0}}
				for _, rrVal := range rectangles[rKey:] {
					if x < rrVal[0] {
						break
					} else if x < rrVal[2] {
						lines = addLine()
					}
				}
			}
		}
	}
}
