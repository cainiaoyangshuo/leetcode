/**
 * @desc 198.打家劫舍
 * @date 2022/7/13
 * @user yangshuo
 */

package _98_打家劫舍
/*
 * 思路：求最大数字和，且不能有相邻的数。和【53.最大子数组和】相似
 * 动态规划：用dp[i]表示前i间能偷到的最高总金额 = 上上个金额 + 当前房子金额 和 上个房子金额的最大值。状态转移方程：
 * dp[i] = max(dp[i-2] + nums[i], dp[i-1])
 * 边界条件：只有一个，只有两个
 */
func rob(nums []int) int {
	n := len(nums)

	//边界条件
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
