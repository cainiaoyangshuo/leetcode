/**
 * @desc 409.最长回文串 https://leetcode-cn.com/problems/longest-palindrome/
 * @date 2022/5/9
 * @user yangshuo
 */

package _09_longestPalindrome

/**
 * 思路：回文数最大长度分为两种情况，1：都为偶数，那么最大回文数为字符串长度；2.有奇数，最大长度为偶数和+1。可以用map维护次数为奇数的字符
 */
func longestPalindrome(s string) int {
	odd := map[int32]int{}
	res := 0

	for _, c := range s {
		if _,ok := odd[c]; ok {
			delete(odd, c)
		} else {
			odd[c] = 1
		}
	}

	//奇数集合为0
	if len(odd) == 0 {
		res = len(s)
	} else {
		res = len(s) - len(odd) + 1
	}

	return res
}
