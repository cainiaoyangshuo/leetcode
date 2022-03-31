package _40_SearchMatrix


func SearchMatrix(matrix [][]int, target int) bool {
	for _, v := range matrix {
		// 对v进行二分查找
		left, right := 0, len(v)-1
		for left <= right {
			mid := (left + right) / 2
			if v[mid] > target {
				right = mid - 1
			} else if v[mid] < target {
				left = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// 优化 O(m+n) 从(0, m-1)开始找
func SearchMatrix01(matrix [][]int, target int) bool {
	cowNum, columnNum := len(matrix), len(matrix[0])
	// 起始点: 右上角
	x, y := 0, columnNum-1

	for x < cowNum && y >= 0 {
		if matrix[x][y] == target {
			return true
		}

		// 不相等判断大小: 目标值小于当前值则在该行往前找，大于当前值则在下一行找
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}

	return false
}
