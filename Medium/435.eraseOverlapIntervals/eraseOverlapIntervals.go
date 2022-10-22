/**
 * @desc 435.无重叠区间 https://leetcode-cn.com/problems/non-overlapping-intervals/
 * @date 2022/5/7
 * @user yangshuo
 */

package _35_eraseOverlapIntervals

import "sort"

/**
 * 思路：贪心。求不重叠区间的个数，所有区间按右端点从小到大排序，用right维护上个区间的右端点，遍历区间，若当前区间[li,ri]与上个区间不重叠，即li >= right
 * 那么就选择这个区间，将right更新为ri
 */
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	//右区间从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	ans, right := 1, intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] >= right {
			right = v[1]
			ans++
		}
	}

	return n-ans
}
