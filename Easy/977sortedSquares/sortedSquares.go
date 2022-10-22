/**
 * @desc 977.有序数组的平方 https://leetcode-cn.com/problems/squares-of-a-sorted-array/
 * @date 2022/4/22
 * @user yangshuo
 */

package _77sortedSquares

/**
 时间复杂度:O(n)
 空间复杂度:O(1)
 */
func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	res := make([]int, n)

	//选择大的数，逆序放入数组中，并移动指针；在所有元素大于等于0时，right至多走n-1次。所以不需要考率边界。
	for i := n-1; i >= 0; i-- {
		if pL, pR := nums[left]*nums[left], nums[right]*nums[right]; pL > pR {
			res[i] = pL
			left++
		} else {
			res[i] = pR
			right--
		}
	}

	return res
}
