/**
 * @desc
 * @date 2022/5/17
 * @user yangshuo
 */

package _44_反转字符串

func reverseString(s []byte)  {
	for i,j := 0,len(s); i < j/2; i++ {
		s[i], s[j-1-i] = s[j-1-i], s[i]
	}
}
