package Algo

import (
	"fmt"
	"sort"
)

type cubeSwapper [][]int

func (c cubeSwapper) Len() int {
	return len(c)
}

func (c cubeSwapper) Less(i, j int) bool {
	if c[i][0] != c[j][0] {
		return c[i][0] < c[j][0]
	} else if c[i][1] != c[j][1] {
		return c[i][1] < c[j][1]
	} else {
		return c[i][2] < c[j][2]
	}
}

func (c cubeSwapper) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func setTop(tops *[][]int, node []int) {
	if len(*tops) == 0 {
		*tops = append(*tops, node)
	} else {
		height := 0 // 当前可添加该立体的锥体中，最高的数值
		for _, v := range *tops {
			if node[1] <= v[1] && node[2] <= v[2] {
				if height < v[0] {
					height = v[0]
				}
			}
		}
		node[0] += height
		*tops = append(*tops, node)
	}
}
func MaxHeight(cuboids [][]int) int {
	for k := range cuboids {
		sort.Sort(sort.Reverse(sort.IntSlice(cuboids[k])))
	}
	sort.Sort(sort.Reverse(cubeSwapper(cuboids)))
	fmt.Println(cuboids)
	tops := make([][]int, 0)
	for _, v := range cuboids {
		setTop(&tops, v)
	}
	res := 0
	for _, v := range tops {
		if res < v[0] {
			res = v[0]
		}
	}
	return res
}
