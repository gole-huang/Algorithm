package algo

import (
	"fmt"
	"strings"
)

func crackSafe(n int, k int) string {
	res := strings.Repeat("0", n-1)
	switch n {
	case 1:
		for n1 := k - 1; n1 >= 0; n1-- {
			res += fmt.Sprintf("%d", n1)
		}
	case 2:
		for n2 := k - 1; n2 >= 0; n2-- {
			for n1 := k - 1; n1 > n2; n1-- {
				res += fmt.Sprintf("%d%d", n1, n2)
			}
			res += fmt.Sprintf("%d", n2)
		}
	case 3:
		for n3 := k - 1; n3 >= 0; n3-- {
			for n2 := k - 1; n2 >= n3; n2-- {
				for n1 := k - 1; n1 > n3; n1-- {
					res += fmt.Sprintf("%d%d%d", n1, n2, n3)
				}
			}
			res += fmt.Sprintf("%d", n3)
		}
	case 4:
		for n4 := k - 1; n4 >= 0; n4-- {
			for n3 := k - 1; n3 >= n4; n3-- {
			L3:
				for n2 := k - 1; n2 >= n4; n2-- {
					for n1 := k - 1; n1 > n4; n1-- {
						if n1 == n3 && n2 == n4 {
							res += fmt.Sprintf("%d%d", n3, n4)
							break L3
						} else {
							res += fmt.Sprintf("%d%d%d%d", n1, n2, n3, n4)
						}
					}
				}
			}
			res += fmt.Sprintf("%d", n4)
		}
	}
	return res
}
