package _12sortArray

import (
	"math/rand"
	"time"
)

/**
 * @desc 快排
 * @date 2022/4/18
 * @user YangShuo
 */
// 快排中心思想:
// 选择一个分界点，把数组分为前后两部分，前部分比后部分的数要小。然后再递归调用函数对两部分分别处理，使得最终数组为有序。
// 其中分界点的选择，可分为1取中心点，2随机取。随机取更为合理。
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := []int{}, []int{}
	pivot := nums[0]

	for _, v := range nums {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = sortArray(left)
	right = sortArray(right)
	return append(append(left, pivot), right...)
}


func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var sort func(nums []int, left, right int)
	sort = func(nums []int, left, right int) {
		if left > right {
			return
		}
		pos := randomPartition(nums, left, right)
		sort(nums, left, pos-1)
		sort(nums, pos+1, right)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func swap(nums []int, a,b int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func randomPartition(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(right - left + 1) + left
	swap(nums, randIndex, left)

	// 基准值
	pivot := nums[left]
	lt := left

	// 循环不变量
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			lt++
			swap(nums, i, lt)
		}
	}

	if lt != left {
		swap(nums, left, lt)
	}

	return lt
}