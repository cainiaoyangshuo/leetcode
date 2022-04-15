package _69majorityElement

/*
 * 多数元素
 * https://leetcode-cn.com/problems/majority-element/
 * 摩尔投票法，遇到相同的数，就投一票，遇到不同的数，就减一票，最后还存在票的数就是众数
 */

func MajorityElement(nums []int) int {
	res := -1
	count := 0

	for _, n := range nums {
		if count == 0 {
			res = n
		}

		if res == n {
			count++
		} else {
			count--
		}
	}

	return res
}
