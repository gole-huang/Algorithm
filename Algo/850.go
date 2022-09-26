package algo

import "sort"

type node struct {
	id  int
	x   int
	add bool
}

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

func totalLength(lines [][2]int) (res int) {
	mergeSegment := make([][2]int, 0)
	for _, v := range lines {
		mergeSegment = addSegment(mergeSegment, v[0], v[1])
	}
	for _, v := range mergeSegment {
		res += v[1] - v[0]
	}
	return
}

func RectangleArea(rectangles [][]int) int {
	if len(rectangles) == 1 {
		return (rectangles[0][2] - rectangles[0][0]) * (rectangles[0][3] - rectangles[0][1]) % 1000000007
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
	xAxis := make([]node, len(rectangles)*2)
	for k, v := range rectangles {
		xAxis[k*2].id = k
		xAxis[k*2].x = v[0]
		xAxis[k*2].add = true
		xAxis[k*2+1].id = k
		xAxis[k*2+1].x = v[2]
		xAxis[k*2+1].add = false
	}
	sort.Slice(xAxis, func(i, j int) bool {
		return xAxis[i].x < xAxis[j].x
	})
	res := 0
	squares := make([][2]int, 0)
	for k := range xAxis {
		if xAxis[k].add {
			squares = append(squares, [2]int{rectangles[xAxis[k].id][1], rectangles[xAxis[k].id][3]})
		} else {
			for cur := range squares {
				if squares[cur] == [2]int{rectangles[xAxis[k].id][1], rectangles[xAxis[k].id][3]} {
					squares = append(squares[:cur], squares[cur+1:]...)
					break
				}
			}
		}
		if k < len(xAxis)-1 && xAxis[k+1].x != xAxis[k].x {
			res += (xAxis[k+1].x - xAxis[k].x) * totalLength(squares)
		}
	}
	return res % 1000000007
}
