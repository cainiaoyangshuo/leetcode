/**
 * @desc 查找出现频率前三的单词
 * @date 2022/10/12
 * @user yangshuo
 */

package leetcode

import (
	"sort"
	"strings"
)

func Top3Words(str string) []string {
	arr := strings.Split(str, " ")

	m := map[string]int{}
	type words struct {
		Word string
		Cnt int
	}
	arrCount := make([]words, 0, len(m))

	for _, s := range arr {
		m[s]++
	}

	for k,v := range m {
		arrCount = append(arrCount, words{k,v})
	}

	sort.Slice(arrCount, func(i, j int) bool {
		return arrCount[i].Cnt > arrCount[j].Cnt
	})

	arrCount = arrCount[:3]

	return []string{arrCount[0].Word, arrCount[1].Word, arrCount[2].Word}
}

// Top3Words2 ///
// 桶排序
func Top3Words2(str string) []string {
	res := []string{}
	arr := strings.Split(str, " ")

	m := map[string]int{}

	maxCnt := 0
	for _, s := range arr {
		m[s]++
		maxCnt = max(m[s], maxCnt)
	}

	// 创建桶，1-max个，每个桶里用数组保存单词
	buckets := make([][]string, maxCnt+1)
	for word, cnt := range m {
		buckets[cnt] = append(buckets[cnt], word)
	}

	for i := maxCnt; i > 0; i-- {
		for _, word := range buckets[i] {
			if len(res) > 3 {
				break
			}
			res = append(res, word)
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