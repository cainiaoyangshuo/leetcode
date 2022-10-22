/**
 * @desc
 * @date 2022/10/15
 * @user yangshuo
 */

package 剑指offer_060_出现频率最高的k个数

func topKFrequent(nums []int, k int) []int {
	maxCnt := 0
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		maxCnt = max(m[v], maxCnt)
	}

	buckets := make([][]int, maxCnt+1)
	for num, cnt := range m {
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := []int{}
	for i := maxCnt; i > 0; i-- {
		for _, num := range buckets[i] {
			if len(res) < k {
				res = append(res, num)
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
