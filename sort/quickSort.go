package sort

/**
 * @desc
 * @date 2022/4/19
 * @user yangshuo
 */

func QuickSort(nums []int) []int {
	//快排基于比较，不稳定，平均O(nlogn)，最坏O(n^2)
	//分治思想，选主元，依次将剩余元素的小于主元放其左侧，大放右侧，然后左右两部分同样处理，直至子序列剩余一个元素，结束排序。
	var sort func(nums []int, left, right int) []int
	sort = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}

		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i  < j && nums[j] >= pivot {
				j--
			}

			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}

			// 交换右边第一个小于主元的元素和左边第一个大于主元的元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换
		nums[i], nums[left] = nums[left], nums[i]

		sort(nums, left, i-1)
		sort(nums, i+1, right)
		return nums
	}

	return sort(nums, 0, len(nums)-1)
}
