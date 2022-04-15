package _02removeKdigits
/**
  移除k位数字，成为最小
  https://leetcode-cn.com/problems/remove-k-digits/
  贪心、单调栈
*/

import "strings"


func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	timer := 0
	var remove func(num string, k int) []byte
	remove = func(num string, k int) []byte {
		s := []byte{}
		if timer == k {
			return s
		}
		for i := range num {
			v := num[i]
			if i < len(num)-1 && num[i+1] < v && timer < k {
				timer++
				continue
			}

			s = append(s, v)
		}
		return remove(string(s), k)
	}

	s := remove(num, k)
	res := strings.TrimLeft(string(s), "0")
	if res == "" {
		return "0"
	}

	return res
}


