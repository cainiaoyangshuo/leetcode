/**
 * @desc 290.单词规律 https://leetcode.cn/problems/word-pattern/
 * @date 2022/5/10
 * @user yangshuo
 */

package _90_wordPattern

import "strings"

/**
 * 哈希表, 双映射
 */
func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	n := len(arr)

	if n != len(pattern) {
		return false
	}

	// 两个map维护双映射
	m := map[string]uint8{}
	mP := map[uint8]string{}
	for i, v := range arr {
		p, ok := m[v]
		c, okk := mP[pattern[i]]
		if ok || okk {
			if p != pattern[i] || c != v {
				return false
			}
		} else {
			m[v] = pattern[i]
			mP[pattern[i]] = v
		}
	}

	return true
}
