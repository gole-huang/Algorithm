package algo

/*
该结构主要用于标识某个格四周的感染情况以及防毒墙放置情况

type grid struct {
	left     *grid
	right    *grid
	up       *grid
	down     *grid
	leftBlk  bool
	rightBlk bool
	upBlk    bool
	downBlk  bool
}

type virus struct {
	infected [][]int
	ifMap    map[[2]int]grid
	zones    []*grid
	walls    map[[2]int]bool //防毒墙，表示方式为(x,x+1),(y,y+1)...
}

//zones:找出所有病毒感染区域，
func (v *virus) virusZone() {
	v.zones = make([]*grid, 0)
	infectMap := make(map[[2]int]bool, 0)
	for i := 0; i < len(v.infected); i++ {
		for j := 0; j < len(v.infected[i]); j++ {
			if v.infected[i][j] == 1 && !infectMap[[2]int{i, j}] { //若某一区域被感染，且未被标记
				v.zones = append(v.zones, v.searchZone(i, j, infectMap))
			}
		}
	}
}

func (v *virus) searchZone(x, y int, ifMap map[[2]int]bool) *grid {
	res := new(grid)
	if ifMap[[2]int{x, y}] {
		return nil
	} else {
		(*res).axis = [2]int{x, y}
		ifMap[[2]int{x, y}] = true
	}
	if x > 0 {
		(*res).up = v.searchZone(x-1, y, ifMap)
	}
	if x < len(v.infected)-1 {
		(*res).down = v.searchZone(x+1, y, ifMap)
	}
	if y > 0 {
		(*res).left = v.searchZone(x, y-1, ifMap)
	}
	if y < len(v.infected)-1 {
		(*res).right = v.searchZone(x, y+1, ifMap)
	}
	return res
}

//病毒增长，需规避防毒墙
func (v *virus) virusGrow(virusZone *grid) {
	if virusZone.axis[1] > 0 && !v.walls[[2]int{virusZone.axis[1] - 1, virusZone.axis[1]}] {
		if virusZone.left == nil {
			virusZone.left = new(grid)
			virusZone.left.axis = [2]int{virusZone.axis[0], virusZone.axis[1] - 1}
			v.infected[virusZone.axis[0]][virusZone.axis[1]-1] = 1
		} else {
			v.virusGrow(virusZone.left)
		}
	}
	if virusZone.axis[1] < len(v.infected[0]) && !v.walls[[2]int{virusZone.axis[1], virusZone.axis[1] + 1}] {
		if virusZone.right == nil {
			virusZone.right = new(grid)
			virusZone.right.axis = [2]int{virusZone.axis[0], virusZone.axis[1] + 1}
			v.infected[virusZone.axis[0]][virusZone.axis[1]+1] = 1
		} else {
			v.virusGrow(virusZone.right)
		}
	}
	if virusZone.axis[0] > 0 && !v.walls[[2]int{virusZone.axis[0] - 1, virusZone.axis[0]}] {
		if virusZone.up == nil {
			virusZone.up = new(grid)
			virusZone.up.axis = [2]int{virusZone.axis[0] - 1, virusZone.axis[1]}
			v.infected[virusZone.axis[0]-1][virusZone.axis[1]] = 1
		} else {
			v.virusGrow(virusZone.up)
		}
	}
	if virusZone.axis[0] < len(v.infected) && !v.walls[[2]int{virusZone.axis[0], virusZone.axis[0] + 1}] {
		if virusZone.down == nil {
			virusZone.down = new(grid)
			virusZone.down.axis = [2]int{virusZone.axis[0] + 1, virusZone.axis[1]}
			v.infected[virusZone.axis[0]+1][virusZone.axis[1]] = 1
		} else {
			v.virusGrow(virusZone.down)
		}
	}
}

//建设防毒墙
func (v *virus) buildWalls(zone *grid) int {
	wallCnt := 0
	if zone.axis[1] > 0 && !v.walls[[2]int{zone.axis[1] - 1, zone.axis[1]}] {
		if zone.left == nil {
			v.walls[[2]int{zone.axis[1] - 1, zone.axis[1]}] = true
			wallCnt++
		} else {
			wallCnt += v.buildWalls(zone.left)
		}
	}
	if zone.axis[1] < len(v.infected[0]) && !v.walls[[2]int{zone.axis[1], zone.axis[1] + 1}] {
		if zone.right == nil {
			v.walls[[2]int{zone.axis[1], zone.axis[1] + 1}] = true
			wallCnt++
		} else {
			wallCnt += v.buildWalls(zone.right)
		}
	}
	if zone.axis[0] > 0 && !v.walls[[2]int{zone.axis[0] - 1, zone.axis[0]}] {
		if zone.up == nil {
			v.walls[[2]int{zone.axis[0] - 1, zone.axis[0]}] = true
			wallCnt++
		} else {
			wallCnt += v.buildWalls(zone.up)
		}
	}
	if zone.axis[0] < len(v.infected) && !v.walls[[2]int{zone.axis[0], zone.axis[0] + 1}] {
		if zone.down == nil {
			v.walls[[2]int{zone.axis[0], zone.axis[0] + 1}] = true
			wallCnt++
		} else {
			wallCnt += v.buildWalls(zone.down)
		}
	}
	return wallCnt
}

//求每个区域病毒扩展数量
func (v *virus) virusExtend(zone *grid) int {
	ext := 0
	if zone.axis[1] > 0 && !v.walls[[2]int{zone.axis[1] - 1, zone.axis[1]}] {
		if zone.left == nil {
			ext++
		} else {
			ext += v.virusExtend(zone.left)
		}
	}
	if zone.axis[1] < len(v.infected[0]) && !v.walls[[2]int{zone.axis[1], zone.axis[1] + 1}] {
		if zone.right == nil {
			ext++
		} else {
			ext += v.virusExtend(zone.right)
		}
	}
	if zone.axis[0] > 0 && !v.walls[[2]int{zone.axis[0] - 1, zone.axis[0]}] {
		if zone.up == nil {
			ext++
		} else {
			ext += v.virusExtend(zone.up)
		}
	}
	if zone.axis[0] < len(v.infected) && !v.walls[[2]int{zone.axis[0], zone.axis[0] + 1}] {
		if zone.down == nil {
			ext++
		} else {
			ext += v.virusExtend(zone.down)
		}
	}
	return ext
}

func (v *virus) findMaxExtend() *grid {
	var res *grid
	max := 0
	for _, zone := range v.zones {
		tmp := v.virusExtend(zone)
		if max < tmp {
			max = tmp
			res = zone
		}
	}
	return res
}

func ContainVirus(isInfected [][]int) int {

}
*/
