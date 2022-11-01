package algo

import (
	"math"
	"sort"
)

func BestCoordinate(towers [][]int, radius int) []int {
	sort.Slice(towers, func(i, j int) bool {
		if towers[i][0] == towers[j][0] {
			return towers[i][1] < towers[j][1]
		} else {
			return towers[i][0] < towers[j][0]
		}
	})
	ymin, ymax := 50, 0
	for _, v := range towers {
		if ymax < v[1] {
			ymax = v[1]
		}
		if ymin > v[1] {
			ymin = v[1]
		}
	}

	maxSighal := 0
	res := []int{0, 0}
	for x := towers[0][0]; x <= towers[len(towers)-1][0]; x++ {
		for y := ymin; y <= ymax; y++ {
			curSignal := 0
			for _, v := range towers {
				length := math.Sqrt(float64((x-v[0])*(x-v[0]) + (y-v[1])*(y-v[1])))
				if length <= float64(radius) {
					curSignal += int(float64(v[2]) / (1 + length))
				}
			}
			if maxSighal < curSignal {
				maxSighal = curSignal
				res[0], res[1] = x, y
			}
		}
	}
	return res
}
