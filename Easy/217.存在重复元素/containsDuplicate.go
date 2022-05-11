/**
 * @desc 217.存在重复元素 https://leetcode.cn/problems/contains-duplicate/
 * @date 2022/5/11
 * @user yangshuo
 */

package _17_存在重复元素

func containsDuplicate(nums []int) bool {
	m := map[int]int{}

	for _,v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 1
	}

	return false
}
