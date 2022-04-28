/**
 * @desc 颜色分类 https://leetcode-cn.com/problems/sort-colors/ 双指针
 * @date 2022/4/28
 * @user yangshuo
 */
package _5sortcolors

func sortColors(nums []int)  {
	n := len(nums)
	p0, p1 := 0, 0

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		} else if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]

			//p1表示1后一位的位置(包括0)，如果p0所在位置为1，那么也要将1交换到p1所在位置
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}

			p0++
			p1++
		}
	}
}
