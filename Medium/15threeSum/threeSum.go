/**
 * @desc 15.三数之和 https://leetcode.cn/problems/3sum/
 * @date 2022/4/22
 * @user yangshuo
 */

package _5threeSum

import "sort"

/**
 * 思路： 排序+双指针  三次遍历可以取得所有组合，包括重复的，去重需要用哈希表。为了去重，先进行排序；再用双指针代替第三重遍历，用一个从右到左的指针right
 * 遍历第一层a，同时来寻找和为[-a]的b和c，sum=b+c, sum=-a,分为三种情况
 * 1.sum = -a.找到了，将a、b、c记录到输出数组，并继续将b、c走到下一个不等于它的位置。
 * 2.sum > -a，说明b+c大了，c--
 * 3.sum < -a，b++
 */
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

		//初始化双指针
		left, right := i + 1, len(nums) - 1

		//计算两数之和b、c的和为负的a
		target := -nums[i]

		//遍历a右边的数组，当nums[left]+nums[right]=target时即找到，同时考率到b、c可能存在重复的情况，也要进行去重，即跳到下个不等于b、c的位置。
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

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
