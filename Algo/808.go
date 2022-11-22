package algo

import "math"

func soupServings(n int) float64 {
	if n > 4275 {
		return 1
	} else if n == 0 {
		return 0.5
	}
	k := int(math.Ceil(float64(n) / 25))
	var lg int
	if k < 4 {
		lg = 4
	} else {
		lg = k
	}
	finA, finB := make([][]float64, lg), make([][]float64, lg)
	for i := 0; i < lg; i++ {
		finA[i], finB[i] = make([]float64, lg), make([]float64, lg)
	}
	finA[0][0], finA[0][1], finA[0][2], finA[0][3] = 0.25, 0.5, 0.75, 1
	finA[1][0], finA[1][1], finA[1][2], finA[1][3] = 0.25, 0.5, 0.75, 0.8125
	finA[2][0], finA[2][1], finA[2][2], finA[2][3] = 0.25, 0.5, 0.5625, 0.6875
	finA[3][0], finA[3][1], finA[3][2], finA[3][3] = 0.25, 0.3125, 0.4375, 0.625
	finB[0][0], finB[0][1], finB[0][2], finB[0][3] = 0.0, 0.0, 0.0, 0.0
	finB[1][0], finB[1][1], finB[1][2], finB[1][3] = 0.25, 0.25, 0.25, 0.0
	finB[2][0], finB[2][1], finB[2][2], finB[2][3] = 0.5, 0.5, 0.25, 0.0625
	finB[3][0], finB[3][1], finB[3][2], finB[3][3] = 0.75, 0.5, 0.25, 0.1875
	for i := 0; i < lg; i++ {
		for j := 0; j < lg; j++ {
			if i > 3 || j > 3 {
				switch {
				case i == 0:
					finA[0][j] = 1
					finB[0][j] = 0
				case i == 1:
					finA[1][j] = (3 + finA[0][j-3]) / 4
					finB[1][j] = finB[0][j-3] / 4
				case i == 2:
					finA[i][j] = (2 + finA[0][j-2] + finA[1][j-3]) / 4
					finB[i][j] = (finB[0][j-2] + finB[1][j-3]) / 4
				case i == 3:
					finA[3][j] = (1 + finA[0][j-1] + finA[1][j-2] + finA[2][j-3]) / 4
					finB[3][j] = (finB[0][j-1] + finB[1][j-2] + finB[2][j-3]) / 4
				case j == 0:
					finA[i][0] = finA[i-4][0] / 4
					finB[i][0] = (finB[i-4][0] + 3) / 4
				case j == 1:
					finA[i][1] = (finA[i-4][1] + finA[i-3][0]) / 4
					finB[i][1] = (finB[i-4][1] + finB[i-3][0] + 2) / 4
				case j == 2:
					finA[i][2] = (finA[i-4][2] + finA[i-3][1] + finA[i-2][0]) / 4
					finB[i][2] = (finB[i-4][2] + finB[i-3][1] + finB[i-2][0] + 1) / 4
				default:
					finA[i][j] = (finA[i-4][j] + finA[i-3][j-1] + finA[i-2][j-2] + finA[i-1][j-3]) / 4
					finB[i][j] = (finB[i-4][j] + finB[i-3][j-1] + finB[i-2][j-2] + finB[i-1][j-3]) / 4
				}
			}
		}
	}
	return (finA[k-1][k-1]-finB[k-1][k-1])/2 + 0.5
}
