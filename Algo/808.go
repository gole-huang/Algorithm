package algo

import "math"

func soupServings(n int) float64 {
	if n > 4275 {
		return 1
	}
	k := int(math.Ceil(float64(n) / 25))
	var lg int
	if k < 5 {
		lg = 5
	} else {
		lg = k + 1
	}
	finA, finB := make([][]float64, lg), make([][]float64, lg)
	for i := 0; i < lg; i++ {
		finA[i], finB[i] = make([]float64, lg), make([]float64, lg)
	}
	finA[1][1], finA[1][2], finA[1][3], finA[1][4] = 0.25, 0.5, 0.75, 1
	finA[2][1], finA[2][2], finA[2][3], finA[2][4] = 0.25, 0.5, 0.75, 0.8125
	finA[3][1], finA[3][2], finA[3][3], finA[3][4] = 0.25, 0.5, 0.5625, 0.6875
	finA[4][1], finA[4][2], finA[4][3], finA[4][4] = 0.25, 0.3125, 0.4375, 0.625
	finB[1][1], finB[1][2], finB[1][3], finB[1][4] = 0.0, 0.0, 0.0, 0.0
	finB[2][1], finB[2][2], finB[2][3], finB[2][4] = 0.25, 0.25, 0.25, 0.0
	finB[3][1], finB[3][2], finB[3][3], finB[3][4] = 0.5, 0.5, 0.25, 0.0625
	finB[4][1], finB[4][2], finB[4][3], finB[4][4] = 0.75, 0.5, 0.25, 0.1875
	for i := 1; i < lg; i++ {
		for j := 1; j < lg; j++ {
			if i > 4 || j > 4 {
				switch {
				case i == 1:
					finA[1][j] = 1
					finB[1][j] = 0
				case i == 2:
					finA[2][j] = (3 + finA[1][j-3]) / 4
					finB[2][j] = finB[1][j-3] / 4
				case i == 3:
					finA[3][j] = (2 + finA[1][j-2] + finA[2][j-3]) / 4
					finB[3][j] = (finB[1][j-2] + finB[2][j-3]) / 4
				case i == 4:
					finA[4][j] = (1 + finA[1][j-1] + finA[2][j-2] + finA[3][j-3]) / 4
					finB[4][j] = (finB[1][j-1] + finB[2][j-2] + finB[3][j-3]) / 4
				case j == 1:
					finA[i][1] = finA[i-4][1] / 4
					finB[i][1] = (3 + finB[i-4][1]) / 4
				case j == 2:
					finA[i][2] = (finA[i-4][2] + finA[i-3][1]) / 4
					finB[i][2] = (2 + finB[i-4][2] + finB[i-3][1]) / 4
				case j == 3:
					finA[i][3] = (finA[i-4][3] + finA[i-3][2] + finA[i-2][1]) / 4
					finB[i][3] = (1 + finB[i-4][3] + finB[i-3][2] + finB[i-2][1]) / 4
				default:
					finA[i][j] = (finA[i-4][j] + finA[i-3][j-1] + finA[i-2][j-2] + finA[i-1][j-3]) / 4
					finB[i][j] = (finB[i-4][j] + finB[i-3][j-1] + finB[i-2][j-2] + finB[i-1][j-3]) / 4
				}
			}
		}
	}
	return (finA[k][k]-finB[k][k])/2 + 0.5
}
