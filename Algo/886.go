package algo

func PossibleBipartition(n int, dislikes [][]int) bool {
	grp := make([][2]map[int]bool, 0)
	for _, vd := range dislikes {
		hasGrp := false
		for kg := 0; kg < len(grp); kg++ {
			if grp[kg][0][vd[0]] && grp[kg][0][vd[1]] || grp[kg][1][vd[0]] && grp[kg][1][vd[1]] {
				return false
			} else if grp[kg][0][vd[0]] { //在其中一个grp的左侧发现左值
				hasGrp = true
				noMatch := true
				for ko := range grp {
					if ko == kg {
						continue
					}
					if grp[ko][0][vd[1]] { //在other grp内的左侧找到右值
						for k := range grp[ko][0] { //grp右值与other左值合并
							grp[kg][1][k] = true
						}
						for k := range grp[ko][1] { //grp左值与other右值合并
							grp[kg][0][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					} else if grp[ko][1][vd[1]] { //在other grp内的右侧找到右值
						for k := range grp[ko][0] { //grp左值与other左值合并
							grp[kg][0][k] = true
						}
						for k := range grp[ko][1] { //grp右值与other右值合并
							grp[kg][1][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					}
				}
				if noMatch {
					grp[kg][1][vd[1]] = true //把右值添加到当前的右侧
				}
			} else if grp[kg][1][vd[0]] { //在其中一个grp的右侧发现左值
				hasGrp = true
				noMatch := true
				for ko := range grp {
					if ko == kg {
						continue
					}
					if grp[ko][0][vd[1]] { //在other grp内的左侧找到右值
						for k := range grp[ko][0] { //grp右值与other左值合并
							grp[kg][0][k] = true
						}
						for k := range grp[ko][1] { //grp左值与other右值合并
							grp[kg][1][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					} else if grp[ko][1][vd[1]] { //在other grp内的右侧找到右值
						for k := range grp[ko][1] { //grp左值与other左值合并
							grp[kg][0][k] = true
						}
						for k := range grp[ko][0] {
							grp[kg][1][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						noMatch = false
						break
					}
				}
				if noMatch {
					grp[kg][0][vd[1]] = true //把右值添加到当前的左侧
				}
			} else if grp[kg][0][vd[1]] { //在其中一个grp的左侧发现右值
				hasGrp = true
				noMatch := true
				for ko := range grp {
					if ko == kg {
						continue
					}
					if grp[ko][0][vd[0]] { //在other grp内的左侧找到左值
						for k := range grp[ko][0] {
							grp[kg][1][k] = true
						}
						for k := range grp[ko][1] {
							grp[kg][0][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					} else if grp[ko][1][vd[1]] { //在other grp内的右侧找到右值
						for k := range grp[ko][0] {
							grp[kg][0][k] = true
						}
						for k := range grp[ko][1] {
							grp[kg][1][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					}
				}
				if noMatch {
					grp[kg][1][vd[0]] = true //把左值添加到当前的右侧
				}
			} else if grp[kg][1][vd[1]] { //在其中一个grp的右侧发现右值
				hasGrp = true
				noMatch := true
				for ko := range grp {
					if ko == kg {
						continue
					}
					if grp[ko][0][vd[0]] { //在other grp内的左侧找到左值
						for k := range grp[ko][0] {
							grp[kg][0][k] = true
						}
						for k := range grp[ko][1] {
							grp[kg][1][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					} else if grp[ko][1][vd[0]] { // 在other grp内的右侧找到左值
						for k := range grp[ko][0] {
							grp[kg][1][k] = true
						}
						for k := range grp[ko][1] {
							grp[kg][0][k] = true
						}
						grp = append(grp[:ko], grp[ko+1:]...)
						if ko < kg {
							kg--
						}
						noMatch = false
						break
					}
				}
				if noMatch {
					grp[kg][0][vd[0]] = true //把右值添加到当前的左侧
				}
			}
		}
		if !hasGrp {
			left, right := make(map[int]bool), make(map[int]bool)
			left[vd[0]], right[vd[1]] = true, true
			grp = append(grp, [2]map[int]bool{left, right})
		}
	}
	return true
}
