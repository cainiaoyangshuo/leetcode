/**
 * @desc 415. 字符串相加 https://leetcode-cn.com/problems/add-strings/
 * @date 2022/5/8
 * @user yangshuo
 */

package _15_addStrings

import "strconv"

/**
 * 思路：模拟【竖式加法】，定义两个指针i,j指向num1,num2，定义变量add，维护当前是否有进位，从末尾遍历，如果断的遍历完了，对应位置补0。
 */
func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	res := ""
	add := 0
	i,j := n1-1, n2-1
	for i >= 0 || j >= 0 || add != 0 {
		// 默认都是0，补位，如果遍历完则对应位是0
		var x,y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		sum := x + y + add
		res = strconv.Itoa(sum%10) + res
		add = sum / 10
		i--
		j--
	}

	return res
}

