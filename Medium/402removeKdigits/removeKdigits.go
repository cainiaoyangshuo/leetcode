package _02removeKdigits
/**
  移除k位数字，成为最小
  https://leetcode-cn.com/problems/remove-k-digits/
  贪心、单调栈
*/

import "strings"

//递归 个别case过不了，找不到问题，暂时考率用单调栈实现
func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	timer := 0
	var remove func(num string, k int) string
	remove = func(num string, k int) string {
		s := []byte{}
		if timer == k {
			return num
		}
		for i := range num {
			v := num[i]
			if i < len(num)-1 && num[i+1] < v && timer < k {
				timer++
				continue
			}

			s = append(s, v)
		}
		// 防止死循环 如果最后一个是最大的，则终止递归 移除最后的数字
		if len(s) == len(num) {
			return num[:len(num)-(k-timer)]
		}
		return remove(string(s), k)
	}

	s := remove(num, k)
	res := strings.TrimLeft(s, "0")
	if res == "" {
		return "0"
	}

	return res
}

/*
 * 思路：单调栈，单调递增栈：
 * 如何判断当前字符删不删，受后面的字符影响，所以应该暂存。用什么存，单调栈。为什么是单调栈，因为本题要维护高位递增，要用最后一个数来进行比较，找到右侧第一个比它小的数，用它替换栈顶。
 *
 */
func removeKdigits1(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	stack := []rune{}
	for _, c := range num {
		// 出栈 当前值比栈顶值小，栈顶出栈，当前值入栈
		for k > 0 && len(stack) > 0 && stack[len(stack) -1] > c {
			stack = stack[:len(stack)-1]
			k--
		}

		// 入栈 栈为空且当前值为0时不能入栈
		if c != '0' || len(stack) != 0 {
			stack = append(stack, c)
		}
	}

	// 没删够则继续删 不能直接移除后k个，len(stack) < k时 会越界。
	for len(stack) != 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	res := string(stack)
	if res == "" {
		return "0"
	}

	return res
}