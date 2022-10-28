package algo

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	for km := range obstacleGrid {
		for kn := range obstacleGrid[km] {
			obstacleGrid[km][kn] = 1 - obstacleGrid[km][kn]
		}
	}
	for km := range obstacleGrid {
		for kn := range obstacleGrid[km] {
			if km == 0 {
				if obstacleGrid[0][kn] == 0 {
					for i := kn + 1; i < len(obstacleGrid[0]); i++ {
						obstacleGrid[0][i] = 0
					}
					break
				}
			} else if kn == 0 {
				if obstacleGrid[km][0] == 0 {
					for i := km + 1; i < len(obstacleGrid); i++ {
						obstacleGrid[i][0] = 0
					}
				}
			} else if obstacleGrid[km][kn] == 1 {
				obstacleGrid[km][kn] = obstacleGrid[km-1][kn] + obstacleGrid[km][kn-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
