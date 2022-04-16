package _5searchInsert
/**
   https://leetcode-cn.com/problems/search-insert-position/
   搜索插入位置
   二分
 */
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 没有找到 有两种情况 target比left值大，应该放到left+1的位置；比left小则应该放到left的位置
	if nums[left] < target {
		return left + 1
	}

	return left
}
