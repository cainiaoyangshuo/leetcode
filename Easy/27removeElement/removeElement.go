package _7removeElement
/**
 name:移除元素
 https://leetcode-cn.com/problems/remove-element/
 双指针
 */
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	fast, slow := 0, 0
	for ;fast < length; fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}
