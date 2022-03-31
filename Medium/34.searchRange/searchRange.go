package _4_searchRange


// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	left := searchLeftIndex(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	// 结束位置 = 目标值+1的左侧再-1
	right := searchLeftIndex(nums, target + 1) - 1

	return []int{left, right}
}

/**
目标的左侧边界: 不断左移右侧边界
 */
func searchLeftIndex(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) /2

		//left=right时左侧不变，右侧不断收紧，直到right走到left左侧
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}