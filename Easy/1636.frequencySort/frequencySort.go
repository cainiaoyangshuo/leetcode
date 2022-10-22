/**
 * @desc 1636. 按照频率将数组升序排序 https://leetcode.cn/submissions/detail/373669328/
 * @date 2022/10/16
 * @user yangshuo
 */

package _636_frequencySort

import "sort"

func  frequencySort(nums []int) []int {
	mNumCnt := map[int]int{}

	maxCnt := 0
	for _, num := range nums {
		mNumCnt[num]++
		maxCnt = max(maxCnt, mNumCnt[num])
	}

	buckets := make([][]int, maxCnt+1)
	for num, cnt := range mNumCnt {
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := []int{}
	for _, numset := range buckets {
		if len(numset) == 0 {
			continue
		} else {
			sort.Slice(numset, func(i, j int) bool {
				return numset[i] > numset[j]
			})

			for _, num := range numset {
				for j := 0; j < mNumCnt[num]; j++ {
					res = append(res, num)
				}
			}
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
