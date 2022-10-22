/**
 * @desc 根据字符出现频率排序 https://leetcode.cn/problems/sort-characters-by-frequency/
 * @date 2022/10/22
 * @user yangshuo
 */

package _51_frequencySort

import "strings"

func frequencySort(s string) string {
	m := map[int32]int{}

	for _, c := range s {
		m[c]++
	}

	maxCnt := 0
	for _, cnt := range m {
		maxCnt = max(cnt, maxCnt)
	}

	buckets := make([][]int32, maxCnt+1)
	for c, cnt := range m {
		buckets[cnt] = append(buckets[cnt], c)
	}

	res := ""
	for i := maxCnt; i > 0; i-- {
		for _, c := range buckets[i] {
			res += strings.Repeat(string(c), m[c])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
