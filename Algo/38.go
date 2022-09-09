package algo

import (
	"fmt"
	"strconv"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		var res string
		str := CountAndSay(n - 1)
		char := str[0]
		cnt := 1
		for i := 1; i < len(str); i++ {
			if char == str[i] {
				cnt++
			} else {
				res = res + strconv.Itoa(cnt) + string(char)
				fmt.Printf("%v\n", res)
				char = str[i]
				cnt = 1
			}
		}
		res = res + strconv.Itoa(cnt) + string(char)
		return res
	}
}
