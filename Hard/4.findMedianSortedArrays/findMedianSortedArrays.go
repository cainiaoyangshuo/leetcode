/**
 * @desc
 * @date 2022/5/18
 * @user yangshuo
 */

package __findMedianSortedArrays

/**
 * 1.合并数组 时间复杂度O(m+n)
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := merge(nums1, len(nums1), nums2, len(nums2))
	var res float64
	n := len(arr)
	if n % 2 == 0 {
		res = float64(arr[n/2-1] + arr[n/2]) / 2
	} else {
		res = float64(arr[n/2])
	}

	return res
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1, p2 := 0, 0
	res := []int{}
	for {
		if m == 0 {
			res = append(res, nums2...)
			break
		}
		if n == 0 {
			res = append(res, nums1...)
			break
		}

		if nums1[p1] < nums2[p2] {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}

		if p1 == m && p2 == n {
			break
		} else if p1 == m {
			res = append(res, nums2[p2:]...)
			break
		} else if p2 == n {
			res = append(res, nums1[p1:m]...)
			break
		}
	}

	return res
}

/**
 * 2.二分查找 时间复杂度:O(log(m+n)) 中位数：奇数情况下为(m+n)/2,偶数情况:(m+n)/2-1、(m+n)/2的平均数, 因此可以转化为寻找两个有序数组第k个数
 */
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	return 0
}

/**
 * 双指针
 */
func getK(nums1 []int, nums2 []int, k int) int {
	//
	return 0
}