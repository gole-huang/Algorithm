package Algo

import "fmt"

type parentheses struct {
	left  int
	right int
}

func LongestValidParentheses(s string) int {
	longest := 0
	var res, par []parentheses
	for i := 0; i < len(s); i++ {
		//fmt.Printf("s[%d] = %q\n", i, s[i])
		if s[i] == '(' {
			var tmp parentheses
			tmp.left = i
			par = append(par, tmp)
			fmt.Printf("Push stack, par[%d].left is %d\n", len(par)-1, i)
		} else {
			fmt.Printf("%d in stack.\n", len(par))
			if len(par) > 0 { //只有当左括号入栈，才有计算的可能
				tmp := par[len(par)-1]
				if len(par) > 1 {
					fmt.Printf("Pop stack, par[%d].left is %d\n", len(par)-1, tmp.left)
					par = par[:len(par)-1]
				} else {
					fmt.Printf("Pop stack empty.\n")
					par = make([]parentheses, 0)
				}
				tmp.right = i
				fmt.Printf("tmp.left=%d, tmp.right=%d\n", tmp.left, tmp.right)
				if len(res) > 0 {
					for res[len(res)-1].left > tmp.left && res[len(res)-1].right < tmp.right {
						if len(res) > 1 {
							fmt.Printf("Pop out res[%d] , left=%d, right=%d\n", len(res)-1, res[len(res)-1].left, res[len(res)-1].right)
							res = res[:len(res)-1]
						} else {
							res = make([]parentheses, 0)
							break
						}
					}
				}
				res = append(res, tmp)
			}
		}
	}
	fmt.Printf("res length: %d\n", len(res))
	if len(res) == 0 {
		return 0
	}
	cur := res[0].right - res[0].left + 1
	for i := 1; i < len(res); i++ {
		fmt.Printf("last position is %d , this position is %d\n", res[i-1].right, res[i].left)
		if res[i-1].right+1 == res[i].left {
			fmt.Printf("Continuies\n")
			cur += res[i].right - res[i-1].right
		} else {
			fmt.Printf("Broken,current is %d\n", cur)
			if longest < cur {
				longest = cur
			}
			cur = res[i].right - res[i].left + 1
		}
	}
	if longest < cur {
		longest = cur
	}
	return longest
}
