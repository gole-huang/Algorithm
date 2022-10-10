package algo

import (
	"sort"
)

type numID struct {
	id,num1,num2 int
}

func AdvantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	nums := make([]numID,len(nums2))
	for k,v:=range nums2 {
		nums[k].id=k
		nums[k].num2=v
	}
	sort.Slice(nums, func(i,j int) bool {
		return nums[i].num2 < nums[j].num2
	})
	nouse:=make([]int,0)
	i,j:=0,0
	for k := range nums {
        setVal:=true
		for i<len(nums1) {
		    if nums1[i]>nums[k].num2 {
		    	nums[k].num1 = nums1[i]
                i++
                setVal=false
                break
			} else {
				nouse = append(nouse, nums1[i])
                i++
			}
		}
		if i==len(nums1) && setVal {
			nums[k].num1=nouse[j]
			j++
        }
	}
	sort.Slice(nums, func(i,j int) bool {
		return nums[i].id < nums[j].id
	})
	for k:=range nums {
		nums1[k]=nums[k].num1
	}
	return nums1
}