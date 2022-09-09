package algo

// PreimageSizeFZF /*
// k实际代表了前x个数里一共有多少个5。
func PreimageSizeFZF(k int) int {
	i := k * 4 / 5 //i以5为基数递增，即i=0、1*5、2*5......
	m := 0         //m是i可能包含的5的个数。例如i=5，k=1；i=25，k=6；i=125；k=25+5+1=31
	for i+m < k {
		i++
		m = 0
		for j := i / 5; j > 0; j /= 5 {
			m += j
		}
	}
	if i+m == k {
		return 5
	} else {
		return 0
	}
}
