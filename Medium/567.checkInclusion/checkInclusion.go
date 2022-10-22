/**
 * @desc 567.字符串的排列 https://leetcode.cn/problems/permutation-in-string/
 * @date 2022/5/28
 * @user yangshuo
 */

package _67_checkInclusion

/**
 * 滑动窗口: 两个数组cnt1,cnt2，cnt1统计s1中各字符的个数，cnt2统计当前遍历的子串中字符的个数，
 */
func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	//定义两个数组cnt1, cnt2
	var cnt1, cnt2 [26]int
	for i := 0; i < n1; i++ {
		 cnt1[s1[i] - 'a']++
		 cnt2[s2[i] - 'a']++
	}

	//正序
	if cnt1 == cnt2 {
		return true
	}

	for i := n1; i < n2; i++ {
		cnt2[s2[i] - 'a']++
		cnt2[s2[i - n1] - 'a']--
		if cnt1 == cnt2 {
			return true
		}
	}

	return false
}
