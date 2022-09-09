package algo

func findChar(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func MinDistance(src, dst string) int {
	if len(src) > len(dst) {
		src, dst = dst, src
	}
	if src == "" {
		return len(dst)
	}
	if len(src) == 1 {
		if findChar(dst, src[0]) == -1 {
			return len(dst)
		}
		return len(dst) - 1
	}
	if dst[0] == src[0] {
		return MinDistance(dst[1:], src[1:])
	}
	min := len(dst)
	if len(src) > len(dst) {
		min = len(src)
	}
	// 匹配首字符
	v := findChar(dst, src[0])
	if v != -1 {
		if min > MinDistance(src[1:], dst[v+1:])+v {
			min = MinDistance(src[1:], dst[v+1:]) + v
		}
	}
	// 删除首字符
	if src[1] == dst[0] && min > MinDistance(src[2:], dst[1:])+1 {
		min = MinDistance(src[2:], dst[1:]) + 1
	}
	// 更改
	if min > MinDistance(src[1:], dst[1:])+1 {
		min = MinDistance(src[1:], dst[1:]) + 1
	}
	return min
}
