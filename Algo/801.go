package algo

func MinSwap(nums1 []int, nums2 []int) int {
	cnt:=0
	for i:=1;i<len(nums1);i++ {
		if nums1[i-1] >= nums1[i] || nums2[i-1] >= nums2[i] {
			exc, noexc := 1, 1
			for ;exc<len(nums1)-2-i; exc++ {
				if nums1[i+exc]<nums1[i+exc+1] && nums2[i+exc]<nums2[i+exc+1]{
					break
				}
			}
			for ;noexc<i; noexc++ {
				if nums1[i-noexc-1]<nums2[i-noexc] && nums2[i-noexc-1]<nums1[i-noexc]{
					break
				}
			}
			if noexc<exc {
				cnt+=noexc
			} else {
				cnt+=exc
				i+=exc
			}
		}
	}
	return cnt
}
