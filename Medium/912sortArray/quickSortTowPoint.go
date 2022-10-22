package _12sortArray

import (
	"math/rand"
	"time"
)

/**
 * @desc 快排 双指针
 * @date 2022/4/19
 * @user yangshuo
 */


func QuickSort2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var sort func(nums []int, left, right int)
	sort = func(nums []int, left, right int) {
		// 递归终止条件
		if left > right {
			return
		}
		pos := Partition(nums, left, right)
		sort(nums, left, pos-1)
		sort(nums, pos+1, right)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func Partition(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(right - left + 1) + left
	swap(nums, randIndex, left)
	//基准值
	pivot := nums[left]
	lt, gt := left+1, right

	for {
		// 从左往右找第一个小于基准值的元素
		for lt <= right && nums[lt] < pivot {
			lt++
		}

		// 从后往前找第一个大于基准值的元素
		for gt > left && nums[gt] > pivot {
			gt--
		}

		if lt >= gt {
			break
		}

		// 交换两个元素
		swap(nums, lt, gt)

		// 两个指针继续找，直到重合
		lt++
		gt--
	}

	//// 这步的作用？
	swap(nums, left, gt)
	return gt
}