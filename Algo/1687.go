package algo

import "fmt"

func weight(t [][]int) (sum int) {
	for _, v := range t {
		sum += v[1]
	}
	return
}

func times(t [][][]int) int {
	sum := len(t)
	for k := range t {
		sum++
		for i := 1; i < len(t[k]); i++ {
			if t[k][i][0] != t[k][i-1][0] {
				sum++
			}
		}
	}
	return sum
}

func minimal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	trucks := make([][][]int, 1)
	trucks[0] = make([][]int, 0)
	for k := 0; k < len(boxes); k++ {
		// 当前车能装下第k个盒子
		if boxes[k][1]+weight(trucks[len(trucks)-1]) <= maxWeight && len(trucks[len(trucks)-1]) < maxBoxes {
			// 假设该盒子是最后一个，或者与后面盒子不同目的地
			if k == len(boxes)-1 || boxes[k][0] != boxes[k+1][0] {
				fmt.Println("append last trucks: ", trucks)
				trucks[len(trucks)-1] = append(trucks[len(trucks)-1], boxes[k])
			} else {
				// 与后面的箱子目的地一致
				i, till, tmpWeight := 0, -1, 0
				for i < len(boxes)-k && boxes[k+i][0] == boxes[k][0] {
					if i+1+len(trucks[len(trucks)-1]) <= maxBoxes && tmpWeight+boxes[k+i][1]+weight(trucks[len(trucks)-1]) <= maxWeight {
						till = i
					}
					tmpWeight += boxes[k+i][1]
					i++
				}
				switch {
				case till == 0:
					trucks = append(trucks, [][]int{boxes[k]})
				case till == i-1:
					trucks[len(trucks)-1] = append(trucks[len(trucks)-1], boxes[k:k+i]...)
					k += i
				default:
					load := times(trucks) + 1 + boxDelivering(boxes[k+till+1:], portsCount, maxBoxes, maxWeight)
					unload := times(trucks) + boxDelivering(boxes[k:], portsCount, maxBoxes, maxWeight)
					return minimal(load, unload)
				}
			}
		} else {
			trucks = append(trucks, [][]int{boxes[k]})
		}
	}
	fmt.Println("trucks: ", trucks)
	return times(trucks)
}
