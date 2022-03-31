package _2_trap

func trap(height []int) int {
	res := 0
	length := len(height)
	left, right := 0, length-1
	iLeftMax, jRightMax := 0, 0

	for left < right {
		leftVal := height[left]
		rightVal := height[right]
		iLeftMax = max(iLeftMax, leftVal)
		jRightMax = max(jRightMax, rightVal)

		if iLeftMax < jRightMax {
			res += iLeftMax - leftVal
			left++
		} else {
			res += jRightMax - rightVal
			right--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func trapDP(height []int) int {
//
//}