/**
 * @desc 334.递增的三元子序列 https://leetcode-cn.com/problems/increasing-triplet-subsequence/
 * @date 2022/5/9
 * @user yangshuo
 */

package _34_increasingTriplet

import "math"

/**
 * 思路：贪心，递增子序列，非连续，用两个变量first,second维护子序列前两个元素，first初始值为nums[0]，从nums[1]开始遍历，num = nums[i]
 * 操作：1.如果 num > second, 则找到了，返回true；2.如果 first < num < second，则second的值为num；3. 否则，first值为num
 */
func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := math.MaxInt32

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
