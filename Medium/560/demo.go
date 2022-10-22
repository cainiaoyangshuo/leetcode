/**
 * @desc 560.和为k的子数组 https://leetcode.cn/problems/subarray-sum-equals-k/
 * @date 2022/5/9
 * @user yangshuo
 */

package _60

/**
 * 思路：前缀和+哈希表，哈希表k为前缀和，v为次数，遍历数组时，先计算i的前缀和pre=pre+nums[i]，然后判断pre-k是否在map中，如果在，说明i加上之前的元素能得到k。
 * 因为pre-k会等于0，所有map初始值应为map[0]=1, 最后将i的前缀和放到map
 */
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		// 记录
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}

	return count
}
