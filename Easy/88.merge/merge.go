/**
 * @desc 88.合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/
 * @date 2022/5/12
 * @user yangshuo
 */

package _8_merge

/**
 * 双指针, 定义两个变量p1,p2，初始为0，
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1, p2 := 0, 0
	res := []int{}
	for {
		if m == 0 {
			res = append(res, nums2...)
			break
		}
		if n == 0 {
			res = append(res, nums1...)
			break
		}
		if nums1[p1] < nums2[p2] {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}

		if p1 == m && p2 == n {
			break
		} else if p1 == m {
			res = append(res, nums2[p2:]...)
			break
		} else if p2 == n {
			res = append(res, nums1[p1:m]...)
			break
		}
	}

	copy(nums1, res)
	println(nums1)
}
