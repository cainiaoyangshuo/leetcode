package _3_最大子序和


/**
   最大子数组和
   https://leetcode-cn.com/problems/maximum-subarray/
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