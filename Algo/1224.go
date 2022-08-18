package Algo

func checkLess(less map[int]int) bool {
	onlyOne := true
	for _, v := range less {
		if v != 1 {
			onlyOne = false
		}
	}
	return len(less) == 1 && onlyOne
}

func MaxEqualFreq(nums []int) int {
	res := 0
	equal := 1                   // 当前每个元素需要满足的个数
	more := [2]int{0, 0}         // 比当前满足个数多1，暂存在more里
	equals := make(map[int]bool) // 当前满足个数的元素
	less := make(map[int]int)    // 当前未满足个数的元素及其数量
	for k, num := range nums {
		if num == more[0] {
			for v := range equals {
				less[v] = equal
			}
			equals = make(map[int]bool) // 清空equals
			equal++
			more[1]++
		} else if _, isEqual := equals[num]; isEqual {
			if more[0] != 0 {
				delete(equals, num)
				for v := range equals {
					less[v] = equal
				}
				equals = make(map[int]bool) // 清空equals
				equal++
				equals[num] = true
				equals[more[0]] = true
				more = [2]int{0, 0}
			} else {
				more = [2]int{num, equal + 1}
				delete(equals, num)
			}
		} else if _, isLess := less[num]; isLess {
			less[num]++
			if less[num] == equal {
				equals[num] = true
				delete(less, num)
			}
		} else {
			if equal != 1 {
				less[num] = 1
			} else {
				equals[num] = true
			}
		}
		if equal == 1 {
			res = k
		} else if more[0] != 0 {
			if len(less) == 0 || checkLess(less) && len(equals) == 0 {
				res = k
			}
		} else {
			if checkLess(less) && len(equals) != 0 {
				res = k
			}
		}
	}
	return res + 1
}
