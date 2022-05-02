/**
 * @desc 59.螺旋矩阵② https://leetcode-cn.com/problems/spiral-matrix-ii/
 * @date 2022/5/2
 * @user yangshuo
 */
package _9generateMatrix

func generateMatrix(n int) [][]int {
	left, top, right, bottom := 0, 0, n-1, n-1
	k := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for k <= n*n {
		//填充左到右
		for i := left; i <= right; i++ {
			res[top][i] = k
			k++
		}
		top++

		//填充上到下
		for i := top; i <= bottom; i++ {
			res[i][right] = k
			k++
		}
		right--

		//填充右到左
		for i := right; i >= left; i-- {
			res[bottom][i] = k
			k++
		}
		bottom--

		//填充下到上
		for i := bottom; i >= top; i-- {
			res[i][left] = k
			k++

		}
		left++
	}

	return res
}
