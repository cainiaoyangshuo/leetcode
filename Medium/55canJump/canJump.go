package _5canJump

/*
  跳跃游戏
  https://leetcode-cn.com/problems/jump-game/
  贪心、动态规划
  */

// 从后往前
func canJump(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return true
	}

	idx := length - 1
	// 从后往前，从倒数第二个元素开始，只要值大于等于下标差(等于刚好能走到)，就把当前下标记录下来，如果标记的下标为0，说明能走到终点。
	for i := length - 2; i >= 0; i++ {
		if  nums[i] >= idx - i {
			idx = i
		}
	}

	if idx == 0 {
		return true
	}

	return false
}

// 从前往后
func canJumpI(nums []int) bool {
	mostFar := 0 // 可跳跃最大距离
	n := len(nums)
	if n < 2 {
		return true
	}

	for i := 0; i < n-1; i++ {
		// 需判断能不能走到下一个
		if mostFar >= i {
			mostFar = max(mostFar, i + nums[i])
			if mostFar >= n - 1 {
				return true
			}
		}
	}
	return false
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}