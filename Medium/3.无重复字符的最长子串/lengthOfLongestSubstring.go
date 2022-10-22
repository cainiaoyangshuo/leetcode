/**
 * @desc 3. 无重复字符的最长子串
 * @date 2022/5/26
 * @user yangshuo
 */

package __无重复字符的最长子串

/**
 * 滑动窗口+哈希表
 *
 */
func lengthOfLongestSubstring(s string) int {
	//判断重复
	m := map[byte]int{}
	max := 0
	//右指针，初始值为0，相当于在字符串最左侧，还没开始移动
	r := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		//移动右指针，若出现过则从下一个开始算
		for r < len(s) - 1 && m[s[r+1]] == 0 {
			m[s[r+1]]++
			r++
		}

		//更新最大值
		max = Max(max, r+1 - i)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}