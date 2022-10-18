package algo

import "math"

func AtMostNGivenDigitSet(digits []string, n int) int {
	digs := make([]int, len(digits))
	for k, v := range digits {
		digs[k] = int(v[0] - '0')
	}
	num := make([]int, 0)
	for n != 0 {
		num = append(num, n%10)
		n /= 10
	}
	res := 0
	for d := len(num) - 1; d >= 0; d-- {
        bigger:=true
		for k, v := range digs {
            if num[d]<v {
                if d==0 {
                    res+=k
                } else {
                    res += (k+1) * int(math.Pow(float64(len(digs)), float64(d)))
                    for d>1 {
                        d--
                        res += int(math.Pow(float64(len(digs)), float64(d)))
                    }
                }
                return res
            } else if num[d]==v {
                bigger=false
                res += (k+1) * int(math.Pow(float64(len(digs)), float64(d)))
                break
            }
		}
        if bigger {
            for d>=0 {
                res += int(math.Pow(float64(len(digs)), float64(d)+1))
                d--
            }
            return res
        }
	}
	return res
}