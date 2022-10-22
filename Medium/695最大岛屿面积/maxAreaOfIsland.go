/**
 * @desc 695.岛屿数量 https://leetcode.cn/problems/max-area-of-island/
 * @date 2022/5/30
 * @user yangshuo
 */

package _95最大岛屿面积

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

/**
 * 深度优先搜索
 */
func maxAreaOfIsland(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area = max(dfs(grid, i, j), area)
		}
	}

	return area
}

/**
 * 找当前位置能够连接到的岛屿的个数。
 */
func dfs(grid [][]int, i, j int) (n int) {
	if grid[i][j] == 1 {
		//将经过的点置为0，避免重复扫描
		grid[i][j] = 0
		n++

		//扫描4个方向一直找，直到四个方向都不为1，返回最大岛屿个数。
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]

			//递归终止条件，当前值已经在最边上时，即上下左右超出grid范围
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
				n += dfs(grid, x, y)
			}
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}