/**
 * @desc
 * @date 2022/11/7
 * @user yangshuo
 */

package leetcode

func QuickSort(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	var partition func(nums []int, left, right int) []int
	partition = func(nums []int, left, right int) []int {
		if left > right {
			return nil
		}
		i,j := left, right

		pivot := nums[left]

		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}

			for i < j && nums[i] <= pivot {
				i++
			}

			nums[i], nums[j] = nums[j], nums[i]
		}

		nums[i], nums[left] = nums[left], nums[i]

		partition(nums, left, i-1)
		partition(nums, i+1, right)
		return nums
	}

	return partition(nums, 0, len(nums)-1)
}