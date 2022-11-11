package algo

func pathCost(grid []string, locks map[byte]bool, start, end [2]int) int {
	edge := [][2]int{start}
	dis := 0
	tmpLen, edgeLen, inner := 0, len(edge), 0
	for tmpLen != edgeLen {
		for _, v := range edge[tmpLen:] {
			if v == end {
				return dis
			}
			if v[0] > 0 {
				left := grid[v[0]-1][v[1]]
				if left != '#' && (left < 'A' || 'F' < left || locks[left]) {
					noRep := true
					for _, ev := range edge[inner:] {
						if [2]int{v[0] - 1, v[1]} == ev {
							noRep = false
							break
						}
					}
					if noRep {
						edge = append(edge, [2]int{v[0] - 1, v[1]})
					}
				}
			}
			if v[0] < len(grid)-1 {
				right := grid[v[0]+1][v[1]]
				if right != '#' && (right < 'A' || 'F' < right || locks[right]) {
					noRep := true
					for _, ev := range edge[inner:] {
						if [2]int{v[0] + 1, v[1]} == ev {
							noRep = false
							break
						}
					}
					if noRep {
						edge = append(edge, [2]int{v[0] + 1, v[1]})
					}

				}
			}
			if v[1] > 0 {
				up := grid[v[0]][v[1]-1]
				if up != '#' && (up < 'A' || 'F' < up || locks[up]) {
					noRep := true
					for _, ev := range edge[inner:] {
						if [2]int{v[0], v[1] - 1} == ev {
							noRep = false
							break
						}
					}
					if noRep {
						edge = append(edge, [2]int{v[0], v[1] - 1})
					}
				}
			}
			if v[1] < len(grid[v[0]])-1 {
				down := grid[v[0]][v[1]+1]
				if down != '#' && (down < 'A' || 'F' < down || locks[down]) {
					noRep := true
					for _, ev := range edge[inner:] {
						if [2]int{v[0], v[1] + 1} == ev {
							noRep = false
							break
						}
					}
					if noRep {
						edge = append(edge, [2]int{v[0], v[1] + 1})
					}
				}
			}
		}
		inner = tmpLen
		tmpLen = edgeLen
		edgeLen = len(edge)
		dis++
	}
	if tmpLen == edgeLen {
		return -1
	} else {
		return dis
	}
}

func nextMove(grid []string, locks map[byte]bool, keys [][2]int, start [2]int) int {
	if len(keys) == 1 {
		return pathCost(grid, locks, start, keys[0])
	}
	min := 901
	for k, v := range keys {
		leftKeys := make([][2]int, len(keys)-1)
		copy(leftKeys[:k], keys[:k])
		copy(leftKeys[k:], keys[k+1:])
		openLocks := make(map[byte]bool)
		for lk, lv := range locks {
			openLocks[lk] = lv
		}
		openLocks[grid[v[0]][v[1]]+32] = true
		cost := pathCost(grid, locks, start, v)
		next := nextMove(grid, openLocks, leftKeys, v)
		if cost < 0 || next < 0 {
			continue
		}
		curPath := cost + next
		if min > curPath {
			min = curPath
		}
	}
	if min == 901 {
		return -1
	} else {
		return min
	}
}

func ShortestPathAllKeys(grid []string) int {
	keys := make([][2]int, 0)
	var start [2]int
	for x := range grid {
		for y, v := range grid[x] {
			if 'a' <= v && v <= 'f' {
				keys = append(keys, [2]int{x, y})
			}
			if v == '@' {
				start = [2]int{x, y}
			}
		}
	}
	return nextMove(grid, make(map[byte]bool), keys, start)
}
