package algo

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	leftLow, leftHigh, rightLow, rightHigh := 0, nums[0], 0, 0
	sum, maxNeg, allNeg, setLeft, reachPeek := 0, nums[0], true, true, false
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > maxNeg {
			maxNeg = nums[i]
		}
		if nums[i] > 0 {
			allNeg = false
			if sum >= leftHigh || setLeft {
				setLeft = false
				leftHigh = sum
				if leftLow > rightLow {
					leftLow = rightLow
				}
				rightHigh = rightLow
			} else if sum > rightHigh {
				rightHigh = sum
				reachPeek = true
			}
		} else if sum < rightLow {
			if reachPeek {
				reachPeek = false
				if rightHigh-rightLow > leftHigh-leftLow {
					leftHigh = rightHigh
					leftLow = rightLow
				}
			}
			rightLow = sum
			rightHigh = rightLow
		}
	}
	if allNeg {
		return maxNeg
	}
	if rightHigh-rightLow > leftHigh-leftLow {
		return rightHigh - rightLow
	} else {
		return leftHigh - leftLow
	}
}
