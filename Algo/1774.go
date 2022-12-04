package algo

import "sort"

func ClosestCost(baseCosts []int, toppingCosts []int, target int) int {
	sort.Ints(baseCosts)
	if baseCosts[0] >= target {
		return baseCosts[0]
	}
	sort.Ints(toppingCosts)
	maxBase := len(baseCosts)
	for k := range baseCosts {
		if baseCosts[k] == target {
			return target
		} else if baseCosts[k] > target {
			maxBase = k
			break
		}
	}
	var diff int
	var less bool
	if maxBase == len(baseCosts) {
		diff = target - baseCosts[maxBase-1]
		less = true
	} else {
		diff = baseCosts[maxBase] - target
		less = false
	}
	for i := 0; i < maxBase; i++ {
		topTar := target - baseCosts[i] // 尝试每个baseCosts
		maxTop := len(toppingCosts)
		for k := range toppingCosts {
			if toppingCosts[k] == topTar {
				return target
			} else if toppingCosts[k] > topTar {
				maxTop = k
				break
			}
		}
		if maxTop == 0 && toppingCosts[0]-topTar < diff {
			diff = toppingCosts[0] - topTar
			less = false
		} else if maxTop != len(toppingCosts) && toppingCosts[maxTop]-topTar < diff {
			diff = toppingCosts[maxTop] - topTar
			less = false
		}
		topSum := []int{0}
		for k := range toppingCosts[:maxTop] {
			add1, add2 := make([]int, len(topSum)), make([]int, len(topSum))
			copy(add1, topSum)
			copy(add2, topSum)
			for j := range add1 {
				add1[j] += toppingCosts[k]
				add2[j] += toppingCosts[k] << 1
			}
			topSum = append(topSum, add1...)
			topSum = append(topSum, add2...)
		}
		for k := range topSum {
			if topTar == topSum[k] {
				return target
			} else {
				if topSum[k] > topTar && topSum[k]-topTar < diff {
					diff = topSum[k] - topTar
					less = false
				} else if topSum[k] < topTar && topTar-topSum[k] <= diff {
					diff = topTar - topSum[k]
					less = true
				}
			}
		}
		if topSum[len(topSum)-1] < topTar && topTar-topSum[len(topSum)-1] <= diff {
			diff = topTar - topSum[len(topSum)-1]
			less = true
		}
	}
	if less {
		return target - diff
	}
	return target + diff
}
