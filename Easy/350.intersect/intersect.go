/**
 * @desc 206.两个数组的交集② https://leetcode.cn/problems/intersection-of-two-arrays-ii/
 * @date 2022/5/22
 * @user yangshuo
 */

package _50_intersect

import "sort"

/**
 * 排序+双指针
 */
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := []int{}
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++

		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}

	return res
}

/**
 * 哈希表
 */