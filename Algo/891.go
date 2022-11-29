package algo

import (
	"sort"
)

func SumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	res := 0
	lg := len(nums)
	ten := [...]int{1, 1024, 1048576, 73741817, 511620083, 898961331, 536396504, 270016253, 496641140, 560523804}
	hundred := [...]int{1, 976371285, 499445072, 322050759, 198967538, 390483007, 646459251, 180717820, 900664884, 989772679}
	thousand := [...]int{1, 688423210, 749218516, 165700476, 787943596, 860192575, 554430047, 4413300, 131425474, 53517118}
	tt := [...]int{1, 905611805, 614428880, 165890099, 935394485, 18165451, 753593133, 634495241, 443675373, 563994701}
	pow := func(num int) int {
		res := 1
		if num >= 10000 {
			res = tt[num/10000]
			num %= 10000
		}
		if num >= 1000 {
			res = (res * thousand[num/1000]) % 1000000007
			num %= 1000
		}
		if num >= 100 {
			res = (res * hundred[num/100]) % 1000000007
			num %= 100
		}
		if num >= 10 {
			res = (res * ten[num/10]) % 1000000007
			num %= 10
		}
		if num > 0 {
			res = (res * 1 << num) % 1000000007
		}
		return res
	}
	for i := 0; i < lg/2; i++ {
		pow1 := pow(len(nums) - 1 - i)
		pow2 := pow(i)
		if pow1 < pow2 {
			pow1 += 1000000007
		}
		res += (nums[len(nums)-1-i] - nums[i]) * ((pow1 - pow2) % 1000000007)
	}
	return res % 1000000007
}
