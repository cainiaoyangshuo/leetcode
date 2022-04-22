/**
 * @desc
 * @date 2022/4/22
 * @user yangshuo
 */
package _5threeSum

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})


	for i := 0; i < len(nums); i++ {
		//去重
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		left, right := i + 1, len(nums) - 1
		target := -nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				//left++
				//right--
				//去重
				for left < right {
					left++
					if nums[left] != nums[left - 1] {
						break
					}
				}
				for left < right {
					right--
					if nums[right] != nums[right + 1] {
						break
					}
				}
			} else if sum > target {
				//和大于0了，往小方向走
				right--
			} else {
				//和小于0，往大走
				left++
			}
		}
	}

	return ans
}
