/**
 * @desc 121.买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 * @date 2022/5/22
 * @user yangshuo
 */

package _21_maxProfit

/**
 * 动态规划: 找到最小值后面的最大值的位置，维护一个最小值min和一个最大值差值max, min初始为第一个元素，max初始为0。遍历比较min：
 * 1.大于min，来计算差值，并与原max比较，将较大的更新为max
 * 2.小于min，则更新为min
 * 时间复杂度O(n), 空间O(1)
 */
func maxProfit1(prices []int) int {
	min, max := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			max = Max(max, prices[i] - min)
		} else {
			min = prices[i]
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}