package _6removeDuplicates

/*
 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/solution/shan-chu-pai-xu-shu-zu-zhong-de-zhong-fu-tudo/
 * 双指针
 */
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	slow, fast := 1, 1 //下标为0无需处理
	for ;fast < length; fast++ {
		// fast从下标为1开始遍历，直到遍历到当前值大于前一个，将数组下标为slow的值改为下标为fast的值，同时将slow后移一位。
		if nums[fast] > nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

