/**
 * @desc
 * @date 2022/5/17
 * @user yangshuo
 */

package _57_reverseWords

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := ""

	for _, v := range arr {
		for i := len(v)-1; i >= 0; i-- {
			res += string(v[i])
		}
		res += " "
	}
	s = strings.Trim(res, " ")
	return s
}
