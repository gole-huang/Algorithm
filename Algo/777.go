package algo

func compareGroup(s1,s2 string) bool {
	cntL,cntX,cntR:=0,0,0
    edge,hasX:=false,false
	for i:=0;i<len(s1);i++{ //两次循环
		if cntL<0 || cntR<0 {
			return false
		}
        switch s1[i] {
        case 'L':
            cntL--
        case 'X':
            hasX=true
            cntX++
        case 'R':
            cntR++
        default: return false
        }
        switch s2[i] {
        case 'L':
            cntL++
			edge=true
        case 'X':
            cntX--
        case 'R':
            if edge {
                return false
            }
            cntR--
        default: return false
        }
    }
    if !hasX {
        return s1==s2
    } else {
	    return cntL==0 && cntX==0 && cntR==0
    }
}

func CanTransform(start string, end string) bool {
	if len(start)!=len(end) {
		return false
	} else if len(start)==0 {
		return true
	}
	group:=[]int{0}
    parted:=false
	for i:=0;i<len(start);i++ {
		if start[i]=='L' {
            parted=true
        }else if start[i]=='R' && parted {
            group=append(group,i)
            parted=false
        }
	}
    group=append(group,len(start))
	for i:=1;i<len(group);i++ {
		if !compareGroup(start[group[i-1]:group[i]],end[group[i-1]:group[i]]){
			return false
		}
	}
	return true
}