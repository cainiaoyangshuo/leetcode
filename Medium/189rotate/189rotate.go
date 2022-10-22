/**
 * @desc 189. 轮转数组 https://leetcode-cn.com/problems/rotate-array/
 * @date 2022/4/22
 * @user yangshuo
 */

package _89rotate

/**
额外空间
 */
func rotate(nums []int, k int)  {
	n := len(nums)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[(i+k)%n] = nums[i]
	}

	copy(nums, res)
}

/**
翻转数组
我们以 n=7n=7，k=3k=3 为例进行如下展示：

操作								结果
原始数组							1~2~3~4~5~6~7
翻转所有元素						7~6~5~4~3~2~1
翻转 [0, k mod n - 1] 区间的元素	5~6~7~4~3~2~1
翻转 [k mod n, n - 1] 区间的元素	5~6~7~1~2~3~4

时间复杂度:O(n)
空间复杂度:O(1)
 */
func rotate1(nums []int, k int)  {
	reverse(nums)
	reverse(nums[0:k%len(nums)])
	reverse(nums[k%len(nums):])
}
func reverse(nums []int) {
	for i, j:= 0, len(nums); i < j/2; i++ {
		nums[i], nums[j-1-i] = nums[j-1-i], nums[i]
	}
}
