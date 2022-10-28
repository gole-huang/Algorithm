package algo

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func MinPathSum(grid [][]int) int {
	for km := range grid {
		for kn := range grid[km] {
			if km == 0 && kn == 0 {
				continue
			} else if km == 0 {
				grid[0][kn] += grid[0][kn-1]
			} else if kn == 0 {
				grid[km][0] += grid[km-1][0]
			} else {
				grid[km][kn] += min(grid[km-1][kn], grid[km][kn-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
