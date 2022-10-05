package algo

func ThreeEqualParts(arr []int) []int {
	if len(arr)<3 {
		return []int{-1,-1}
	}
	cnt:=0
	for _,v:=range arr {
		if v==1 {
			cnt++
		}
	}
	if cnt==0 {
		return []int{0,len(arr)-1}
	} else if cnt%3!=0 {
		return []int{-1,-1}
	}
	cnt/=3
	one:=0
	for i:=0;i<len(arr);i++ {
		if arr[i]==0 {
			one++
		} else {
            break
        }
	}
	two,_cnt:=0,cnt+1
	for i:=one;i<len(arr);i++ {
		if arr[i]==1 {
			_cnt--
			if _cnt==0 {
				two=i
				break
			}
		}
	}
	three:=0
	_cnt=cnt+1
	for i:=two;i<len(arr);i++ {
		if arr[i]==1 {
			_cnt--
			if _cnt==0 {
				three=i
				break
			}
		}
	}
	for i:=0;i<len(arr)-three;i++ {
		if arr[three+i]!=arr[one+i] || arr[three+i]!=arr[two+i] {
			return []int{-1,-1}
		}
	}
	return []int{one+len(arr)-three-1, two+len(arr)-three}
}