package _52maxProduct

/**
 最大乘积子数组
 https://leetcode-cn.com/problems/maximum-product-subarray/
 dp
 */
func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	// 最大值，最小值，全局最大值
	max, min, gMax :=nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		// 记录i-1时最大和最小
		mx, mi := max, min
		max = Max(mx * v, Max(v, mi * v))
		min = Min(mi * v, Min(v, mx * v))

		gMax = Max(gMax, max)
	}

	return gMax
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
