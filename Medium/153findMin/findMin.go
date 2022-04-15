package _53findMin

/**
 寻找旋转排序数组中的最小值
 */
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			// if 中点的值大于最右的值，说明最小值在右半部分，左起点从中点后移一位继续二分
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// else 中点值小于最右的话(不存在等于)，说明最小值在左半部分，右起点为中点
			// 为什么不左移呢
			// 答：因为如果中点就是最小值的话会错过
			right = mid
		}
	}
	return nums[left]
	// 为什么最后left为最小值呢？
	// 答：当left=right时循环终止，left和right都可。
}