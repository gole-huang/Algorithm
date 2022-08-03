package Algo

func validPath(x, y int, obstacleGrid [][]int) int {
	if x == len(obstacleGrid)-1 && y == len(obstacleGrid[x])-1 {
		return 1
	} else if x == len(obstacleGrid)-1 {
		if obstacleGrid[x][y+1] == 0 {
			return validPath(x, y+1, obstacleGrid)
		} else {
			return 0
		}
	} else if y == len(obstacleGrid[0])-1 {
		if obstacleGrid[x+1][y] == 0 {
			return validPath(x+1, y, obstacleGrid)
		} else {
			return 0
		}
	} else {
		res := 0
		if obstacleGrid[x][y+1] == 0 {
			res += validPath(x, y+1, obstacleGrid)
		}
		if obstacleGrid[x+1][y] == 0 {
			res += validPath(x+1, y, obstacleGrid)
		}
		return res
	}
}

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	} else {
		return validPath(0, 0, obstacleGrid)
	}
}
