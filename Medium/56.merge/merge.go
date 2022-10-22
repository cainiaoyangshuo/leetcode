/**
 * @desc 56.合并区间  https://leetcode-cn.com/problems/merge-intervals/
 * @date 2022/4/28
 * @user yangshuo
 */

package _6_merge

import "sort"

/**
 * 思路：按区间左边升序，不断通过区间左边来判断是否重叠，重叠则更新右边
 */
func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, v := range intervals {
		left, right := v[0], v[1]

		if len(res) == 0 || res[len(res)-1][1] < left {
			res = append(res, v)
		} else {
			// 如果重叠了，更新数组最后一个区间的最大值
			res[len(res)-1][1] = max(res[len(res)-1][1], right)
		}
	}

	return res
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}