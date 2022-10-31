package algo

import "fmt"

func minSwap(nums1 []int, nums2 []int) int {
	length := len(nums1)
	cnt, back, front := 0, 0, 0
	start := false
	needEx := false
	for i := 1; i < length; i++ {
		switch {
		case nums1[i-1] >= nums1[i] || nums2[i-1] >= nums2[i]:
			if !start {
				start = true
				back++
			}
			needEx = !needEx
			if needEx {
				front++
			} else {
				back++
			}
		case nums1[i-1] >= nums2[i] || nums2[i-1] >= nums1[i]:
			if needEx {
				front++
			} else {
				back++
			}
		default:
			if !start || front == 0 {
				back = 0
				continue
			}
			fmt.Println(front, back)
			if back < front {
				cnt += back
			} else {
				cnt += front
			}
			back, front = 0, 0
			start = false
			needEx = false
		}
	}
	if start {
		if back < front {
			cnt += back
		} else {
			cnt += front
		}
	}
	return cnt
}
