/**
 * @desc 238.除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/
 * @date 2022/5/9
 * @user yangshuo
 */

package _38_productExceptSelf
/**
 * 思路：左右乘积，准对索引i，i左边所有乘积 * i右边所有乘积
 */
func productExceptSelf(nums []int) []int {
	n := len(nums)
	l, r , res := make([]int, n), make([]int, n), make([]int, n)

	// i左侧乘积, 0左侧没有元素，所以l[0]=1
	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = nums[i-1] * l[i-1]
	}

	// i右侧乘积, n-1右侧没有元素，所以r[n-1]=1
	r[n-1] = 1
	for i := n-2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	//
	for i := 0; i < n; i++ {
		res[i] = l[i] * r[i]
	}

	return res
}

/**
 * 空间复杂度优化: O(1) 思路：先用输出数组，记录所有i左侧乘积，再用变量代替数组，维护一个变量r，记录i右侧乘积，从右往左遍历一遍即可。
 */
func productExceptSelf_1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// i左侧乘积, 0左侧没有元素，所以l[0]=1
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	// 用变量R维护i右侧乘积，因最右后面没有元素，初始值为1
	r := 1
	for i := n-1; i >= 0; i-- {
		// i左侧 * i右侧
		res[i] = res[i] * r
		// r为i右侧乘积
		r *= nums[i]
	}

	return res
}
