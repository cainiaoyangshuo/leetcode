/**
 * @desc 283. 移动零 https://leetcode-cn.com/problems/move-zeroes/
 * @date 2022/4/23
 * @user yangshuo
 */
package _83moveZeroes

/**
思路：双指针，左指针指向【已处理好的序列的尾部】，右指针指向【待处理序列的头部】。右指针向右移动，若右指针指向非零数，交换左右指针并左指针右移
性质：
1.左指针左边均为非零数
2.右指针到左指针之间都为0
注：相对顺序，不是指排序
 */
func moveZeroes(nums []int)  {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

/**
计数
思路：不断的将非0数放到第一个0的位置；遍历，若是0，记0出现的次数，若不是0且出现过0，将当前值放到i-k的位置，当前值赋0
 */
func moveZeroes2(nums []int)  {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k++
		} else {
			if k > 0 {
				nums[i-k] = nums[i]
				nums[i] = 0
			}
		}
	}
}