package algo

import "math"

type multi struct {
	mul int
	arr [3]int
}

func pick(res []multi) (int, []multi) {
	min := 0
	for k := range res {
		if res[k].mul < res[min].mul {
			min = k
		}
	}
	ans := res[min].mul
	switch {
	case res[min].arr[0] == 0 && res[min].arr[2] == 0 && res[min].arr[1] > 2:
		tmp := (res[min].arr[1] - 3) / 5
		res[min].mul = res[min].mul / int(math.Pow(5, float64(res[min].arr[1]))) * int(math.Pow(3, float64(tmp*2+1))) * int(math.Pow(7, float64(tmp*3+2)))
		res[min].arr[0] += tmp*2 + 1
		res[min].arr[2] += tmp*3 + 2
		res[min].arr[1] = 0
	case res[min].arr[0] > 1 && res[min].arr[2] > 2:
		res[min].arr[0] -= 2
		res[min].arr[1] += 5
		res[min].arr[2] -= 3
		res[min].mul = res[min].mul / 3087 * 3125
	case res[min].arr[0] > 0 && res[min].arr[2] > 0:
		res[min].arr[0]--
		res[min].arr[2]--
		res[min].arr[1] += 2
		res[min].mul = res[min].mul / 21 * 25
	case res[min].arr[1] == 2 && res[min].arr[2] == 0:
		res = append(res, multi{res[min].mul / 25 * 27, [3]int{res[min].arr[0] + 3, 0, 0}})
		res[min].arr[1]--
		res[min].arr[2]++
		res[min].mul = res[min].mul / 5 * 7
	case res[min].arr[0] == 0 && res[min].arr[1] == 0:
		res = append(res[:min], res[min+1:]...)
	default:
		if res[min].arr[1] > 0 {
			res[min].arr[1]--
			res[min].arr[2]++
			res[min].mul = res[min].mul / 5 * 7
		} else {
			res[min].arr[0]--
			res[min].arr[1]++
			res[min].mul = res[min].mul / 3 * 5
		}
	}
	return ans, res
}

func GetKthMagicNumber(k int) int {
	if k < 5 {
		return k*2 - 1
	}
	nums := make([]int, k)
	nums[0], nums[1], nums[2], nums[3] = 1, 3, 5, 7
	res := []multi{{9, [3]int{2, 0, 0}}}
	for i := 4; i < len(nums); i++ {
		nums[i], res = pick(res)
	}
	return nums[k-1]
}

/*case (res[min].arr[1]-3)%5 == 0:
tmp := (res[min].arr[1] - 3) / 5
res[min].mul = res[min].mul / int(math.Pow(5, float64(res[min].arr[1]))) * int(math.Pow(3, float64(tmp*2+1))) * int(math.Pow(7, float64(tmp*3+2)))
res[min].arr[0] += tmp*2 + 1
res[min].arr[2] += tmp*3 + 2
res[min].arr[1] = 0*/
