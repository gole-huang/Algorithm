package algo

func GroupThePeople(groupSizes []int) [][]int {
	groupMap := make(map[int][]int)
	res := make([][]int, 0)
	for k, v := range groupSizes {
		if _, ok := groupMap[v]; ok {
			groupMap[v] = append(groupMap[v], k)
			if len(groupMap[v]) == v {
				res = append(res, groupMap[v])
				delete(groupMap, v)
			}
		} else {
			if v == 1 {
				res = append(res, []int{k})
			} else {
				groupMap[v] = make([]int, 1)
				groupMap[v][0] = k
			}
		}
	}
	return res
}
