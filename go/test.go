/**
 * @desc 牵手面试：最长无重复子串
 * @date 2022/11/3
 * @user yangshuo
 */

package _go

func getMaxStr(str string) string {

	res := ""
	m := map[byte]int{}

	//右指针
	r := 0
	for i := 0; i < len(str); i++ {
		if i != 0 {
			delete(m, str[i])
		}
		//移动右指针，直到遇到重复的字符。
		for r < len(str) && m[str[r]] == 0 {
			m[str[r]]++
			r++
		}

		//此时i-r之间为非重复子串
		res = max(res, str[i:r])
	}

	return res
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}