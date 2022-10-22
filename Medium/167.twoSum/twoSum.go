/**
 * @desc 167.两数之和②-有序数组 https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
 * @date 2022/5/26
 * @user yangshuo
 */

package _67_twoSum

/**
 * 二分: 时间复杂度O(nlog), 空间O(1)
 */
func twoSum(numbers []int, target int) []int {

	for i, v := range numbers {
		left, right := i + 1, len(numbers) - 1
		n := target - v
		for left <= right {
			mid := (left + right) / 2
			if numbers[mid] == n {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > n {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return []int{}
}

/**
 * 双指针: 时间O(n) 缩小查找范围
 * 双指针left, right分别指向首尾，如果和小于target则left++, 反之，right--，等于则找到。
 * 不存在漏掉答案
 */
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}