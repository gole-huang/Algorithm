package algo

import "fmt"

func calStep(queue []int, total int) (step int) {
	if len(queue) == 1 {
		return 0
	} else if len(queue) == 3 {
		if queue[0] < queue[2] {
			return queue[0] * queue[1]
		}
		return queue[1] * queue[2]
	}
	i, left := 0, 0
	lEdge, rEdge := total>>1, (total-1)>>1
	for left+queue[i] <= lEdge {
		left += queue[i]
		step += left * queue[i+1]
		i += 2
	}

	d, right := len(queue)-1, 0
	for right+queue[d] <= rEdge {
		right += queue[d]
		step += right * queue[d-1]
		d -= 2
	}
	return
}
func minMoves(nums []int, k int) int {
	for nums[0] == 0 {
		nums = nums[1:]
	}
	switch {
	case k == 1:
		return 0
	case k == 2:
		cnt, min := 0, len(nums)-2
		for i := 1; i < len(nums); i++ {
			if nums[i] == 0 {
				cnt++
			} else {
				if min > cnt {
					min = cnt
				}
				cnt = 0
			}
		}
		return min
	case k == len(nums):
		return 0
	}
	queue := []int{1}
	before, cnt1, ptr := true, 1, -1
SetQueue:
	for i := 1; i < len(nums); i++ {
		switch {
		case before && nums[i] == 0:
			queue = append(queue, 1)
			before = false
		case before && nums[i] == 1:
			queue[len(queue)-1]++
			cnt1++
			if cnt1 == k {
				ptr = i + 1
				break SetQueue
			}
		case !before && nums[i] == 0:
			queue[len(queue)-1]++
		case !before && nums[i] == 1:
			queue = append(queue, 1)
			before = true
			cnt1++
			if cnt1 == k {
				ptr = i + 1
				break SetQueue
			}
		}
	}
	curStep := calStep(queue, k)
	minStep, cur0 := curStep, 0
	fmt.Println(minStep)
	for ; ptr != len(nums); ptr++ {
		if nums[ptr] == 0 {
			cur0++
		} else {
			if cur0 == 0 {
				queue[len(queue)-1]++
			} else {
				queue = append(queue, []int{cur0, 1}...)
				cur0 = 0
			}
			queue[0]--
			if queue[0] == 0 {
				queue = queue[2:]
			}
			curStep := calStep(queue, k)
			if minStep > curStep {
				minStep = curStep
			}
		}
	}
	return minStep
}
