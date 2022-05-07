/**
 * @desc 54.螺旋矩阵 https://leetcode-cn.com/problems/spiral-matrix/
 * 思路：按层模拟，左上角(top,left),右上角(bottom.right)，从左到右遍历上侧元素(top,left)到(top, right)；从上到下遍历右侧元素(top+1,right)到(bottom, right);如果left<right且top<bottom, 从右到左遍历下侧元素(bottom,right-1)到(bottom,left+1), 从下到上遍历上侧元素(bottom+1, left)到(top+1, left)。遍历完当前层后，进入到下一层，top.left +1，bottom，right -1、
 * @date 2022/5/3
 * @user yangshuo
 */

package _4spiralOrder

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	left, top, right, bottom := 0, 0, n-1, m-1
	k := 0

	for left <= right && top <= bottom {
		//填充左到右
		for i := left; i <= right; i++ {
			res[k] = matrix[top][i]
			k++
		}


		//填充上到下
		for i := top+1; i <= bottom; i++ {
			res[k] = matrix[i][right]
			k++
		}


		if left < right && top < bottom {
			//填充右到左
			for i := right - 1; i > left; i-- {
				res[k] = matrix[bottom][i]
				k++
			}


			//填充下到上
			for i := bottom; i > top; i-- {
				res[k] = matrix[i][left]
				k++
			}

		}

		//第一圈走完，进入下一圈
		top++
		right--
		bottom--
		left++
	}

	return res
}
