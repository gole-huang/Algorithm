package Algo

func Insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 || newInterval[0] <= intervals[0][0] && newInterval[1] >= intervals[len(intervals)-1][1] {
		return [][]int{newInterval}
	}
	left, right := intervals[0][0], intervals[len(intervals)-1][1]
	if newInterval[1] < left {
		return append([][]int{newInterval}, intervals...)
	} else if newInterval[0] > right {
		return append(intervals, newInterval)
	} else if newInterval[0] < left {
		for i := len(intervals) - 1; i >= 0; i-- {
			if newInterval[1] > intervals[i][1] { //在intervals[i]外部
				return append([][]int{newInterval}, intervals[i+1:]...)
			} else if newInterval[1] >= intervals[i][0] { //在intervals[i]内部
				return append([][]int{{newInterval[0], intervals[i][1]}}, intervals[i+1:]...)
			}
		}
	} else if newInterval[1] > right {
		for i := 0; i < len(intervals); i++ {
			if newInterval[0] < intervals[i][0] { //在intervals[i]外部
				return append(intervals[:i], newInterval)
			} else if newInterval[0] <= intervals[i][1] { //在intervals[i]内部
				return append(intervals[:i], []int{intervals[i][0], newInterval[1]})
			}
		}
	} else {
		l, r := 0, len(intervals)-1
		for ; l < len(intervals); l++ {
			if newInterval[0] < intervals[l][0] { //在intervals[l]外部
				left = newInterval[0]
				break
			} else if newInterval[0] <= intervals[l][1] { //在intervals[l]内部
				left = intervals[l][0]
				break
			}
		}
		for ; r >= 0; r-- {
			if newInterval[1] > intervals[r][1] { //在intervals[r]外部
				right = newInterval[1]
				break
			} else if newInterval[1] >= intervals[r][0] { //在intervals[r]内部
				right = intervals[r][1]
				break
			}
		}
		res := make([][]int, 0)
		res = append(res, intervals[:l]...)
		res = append(res, []int{left, right})
		res = append(res, intervals[r+1:]...)
		return res
	}
	return nil
}
