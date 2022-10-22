/**
  53.最大子数组和 https://leetcode-cn.com/problems/maximum-subarray/
*/

package _3_最大子序和

/**
 * 思路：动态规划。 优化空间 O(1)
 */
func maxSubArray(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	// max记录最大子数组   gMax记录全局最大子数组，初始值为第一个元素
	max, gMax := 0, nums[0]

	for _, v := range nums {
		// 当前最大子数组大于0，则累加；小于0，则重新记录
		if max > 0 {
			max += v
		} else {
			max = v
		}

		// 每次遍历比较全局最大值和前一个最大子数组的最大值，大的为新的全局最大值
		gMax = Max(max, gMax)
	}

	return gMax
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray1(nums []int) int {
	pre, max := 0, nums[0]

	for _, v := range nums {
		pre = Max(pre + v, v)
		max = Max(max, pre)
	}

	return max
}

/**
 * dp 不优化空间, dp数组表示以i结尾的子数组和, 初始值为nums[0], 状态转移情况
 * 1:dp[i-1] > 0, 直接把nums[i]加到dp[i]
 * 2.dp[i-1] <= 0, 把nums[i]加上也不会变大。于是，另起炉灶，dp[i]处即为nums[i]
 */
func maxSubArray2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	//最大和
	max := nums[0]

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		//更新最大和，即当前子数组和和上一个和的较大者。
		max = Max(max, dp[i])
	}

	return max
}