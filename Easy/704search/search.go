package _04search

/**
 * @desc
 * @date 2022/4/20
 * @user yangshuo
 */

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right + left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}